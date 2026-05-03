package tasks

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"github.com/osuTitanic/titanic-go/internal/constants"
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"github.com/osuTitanic/titanic-go/internal/state"
)

// TODO: somehow defer this into cli
var force = true
var rankIndexWorkers = 8

// IndexRanks checks whether leaderboards are empty and rebuilds them when needed.
func IndexRanks(app *state.State, logger *slog.Logger) error {
	topPlayers, err := app.Rankings.TopPlayers(constants.ModeOsu, 0, 1, "performance", nil)
	if err != nil {
		return fmt.Errorf("failed to check leaderboard status: %w", err)
	}
	if len(topPlayers) > 0 && !force {
		logger.Info("Leaderboard is not empty, please clear it first.")
		return nil
	}

	criteria := map[string]any{
		"restricted = ?": false,
		"activated = ?":  true,
	}
	activePlayers, err := app.Repositories.Users.Many(criteria, "Stats")
	if err != nil {
		return fmt.Errorf("failed to fetch active players: %w", err)
	}

	logger.Info(
		"Indexing player ranks...",
		"total_users", len(activePlayers),
		"workers", rankIndexWorkerCount(app, len(activePlayers)),
	)
	return indexRanksForPlayers(app, logger, activePlayers)
}

func indexRanksForPlayer(app *state.State, logger *slog.Logger, player *schemas.User) error {
	country := strings.ToLower(player.Country)

	for _, stats := range player.Stats {
		if err := app.Rankings.Update(stats, country); err != nil {
			return fmt.Errorf("failed to update rankings for user %d mode %d: %w", player.Id, stats.Mode, err)
		}

		if err := app.Rankings.UpdateLeaderScores(stats, country, app.Repositories.Scores); err != nil {
			return fmt.Errorf("failed to update leader rankings for user %d mode %d: %w", player.Id, stats.Mode, err)
		}
	}

	if err := app.Rankings.UpdateKudosu(player.Id, country, app.Repositories.Modding); err != nil {
		return fmt.Errorf("failed to update kudosu rankings for user %d: %w", player.Id, err)
	}

	logger.Info(
		"Updated ranks for player",
		"id", player.Id, "name", player.Name,
	)
	return nil
}

func indexRanksForPlayers(app *state.State, logger *slog.Logger, players []*schemas.User) error {
	workerCount := rankIndexWorkerCount(app, len(players))
	if workerCount == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	jobs := make(chan *schemas.User)
	errs := make(chan error, 1)

	worker := func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case player, ok := <-jobs:
				if !ok {
					return
				}
				if err := indexRanksForPlayer(app, logger, player); err != nil {
					select {
					case errs <- err:
						cancel()
					default:
					}
					return
				}
			}
		}
	}

	for range workerCount {
		wg.Add(1)
		go worker()
	}

	for _, player := range players {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return <-errs
		case jobs <- player:
		}
	}

	close(jobs)
	wg.Wait()

	select {
	case err := <-errs:
		return err
	default:
		return nil
	}
}

func rankIndexWorkerCount(app *state.State, playerCount int) int {
	if playerCount <= 0 {
		return 0
	}
	workerCount := rankIndexWorkers

	if app.Config.PostgresPoolSize > 0 {
		workerCount = app.Config.PostgresPoolSize
	}

	if workerCount < 1 {
		return 1
	}
	if workerCount > playerCount {
		return playerCount
	}
	return workerCount
}
