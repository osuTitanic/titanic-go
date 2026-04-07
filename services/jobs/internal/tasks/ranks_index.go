package tasks

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/osuTitanic/titanic-go/internal/constants"
	"github.com/osuTitanic/titanic-go/internal/state"
)

// TODO: somehow defer this into cli
var force = true

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
	)
	// TODO: Use goroutines to speed up this process

	for _, player := range activePlayers {
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
	}
	return nil
}
