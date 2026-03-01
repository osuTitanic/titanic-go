package state

import (
	"fmt"
	"log/slog"

	"github.com/osuTitanic/common-go/config"
	"github.com/osuTitanic/common-go/database"
	"github.com/osuTitanic/common-go/email"
	"github.com/osuTitanic/common-go/logging"
	"github.com/osuTitanic/common-go/storage"
	"gorm.io/gorm"
)

// State holds the application state, including database
// repositories, configuration, logger, storage & email services.
type State struct {
	*Repositories

	Config   *config.Config
	Database *gorm.DB
	Logger   *slog.Logger
	Storage  storage.Storage
	Email    email.Email
	// TODO: Add redis instance
}

func NewState(environmentFiles ...string) (*State, error) {
	cfg, err := config.LoadConfig(environmentFiles...)
	if err != nil {
		return nil, fmt.Errorf("state: failed to load config: %w", err)
	}

	logLevel := slog.LevelInfo
	if cfg.Debug {
		logLevel = slog.LevelDebug
	}
	logging.SetDefault("titanic", logLevel)
	logger := slog.Default()

	fs := storage.NewFileStorage(cfg.DataPath)
	if err := fs.Setup(); err != nil {
		return nil, fmt.Errorf("state: failed to setup storage: %w", err)
	}

	db, err := database.CreateSession(cfg)
	if err != nil {
		return nil, fmt.Errorf("state: failed to create database session: %w", err)
	}

	mailer, err := email.NewEmailFromConfig(cfg)
	if err != nil {
		database.CloseSession(db)
		return nil, fmt.Errorf("state: failed to create email service: %w", err)
	}

	if err := mailer.Setup(); err != nil {
		database.CloseSession(db)
		return nil, fmt.Errorf("state: failed to setup email service: %w", err)
	}

	return &State{
		Repositories: NewRepositories(db),
		Config:       cfg,
		Database:     db,
		Logger:       logger,
		Storage:      fs,
		Email:        mailer,
	}, nil
}

// DatabaseTransaction executes the given function within a database transaction.
// Example usage:
//
//	err := state.DatabaseTransaction(func(repos *Repositories) error {
//	    // Perform database operations using repos
//	 	repos.User.Create(...)
//
//		// If an error is returned, the transaction will be rolled back
//	    // If nil is returned, the transaction will be committed
//	    return nil
//	})
func (state *State) DatabaseTransaction(fn func(repositories *Repositories) error) error {
	if state == nil || state.Database == nil {
		return fmt.Errorf("state: database is not initialized")
	}
	return state.Database.Transaction(func(tx *gorm.DB) error {
		return fn(NewRepositories(tx))
	})
}

// Close gracefully closes the state
func (state *State) Close() error {
	if state == nil || state.Database == nil {
		return nil
	}
	return database.CloseSession(state.Database)
}
