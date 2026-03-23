package database

import (
	"time"

	"github.com/osuTitanic/common-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CreateSession opens a postgres connection using values from `config.Config`
func CreateSession(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN()), &gorm.Config{
		Logger: NewGormLogger(),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if cfg.PostgresPoolEnabled {
		// Use configured pool values
		sqlDB.SetMaxOpenConns(cfg.PostgresPoolSizeOverflow)
		sqlDB.SetMaxIdleConns(cfg.PostgresPoolSize)
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.PostgresPoolRecycle) * time.Second)
		sqlDB.SetConnMaxIdleTime(time.Duration(cfg.PostgresPoolTimeout) * time.Second)
	}

	if cfg.PostgresPoolPrePing {
		if err := sqlDB.Ping(); err != nil {
			return nil, err
		}
	}
	return db, nil
}

// CloseSession closes the underlying database connection pool
func CloseSession(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
