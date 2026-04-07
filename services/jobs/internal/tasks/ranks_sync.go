package tasks

import (
	"fmt"
	"log/slog"

	"github.com/osuTitanic/titanic-go/internal/state"
)

// UpdateRanks updates rank history for all active users.
func UpdateRanks(app *state.State, logger *slog.Logger) error {
	criteria := map[string]any{
		"restricted = ?": false,
		"activated = ?":  true,
	}
	userList, err := app.Repositories.Users.Many(criteria, "Stats")
	if err != nil {
		return fmt.Errorf("failed to fetch users: %w", err)
	}
	logger.Info("Updating rank history...", "total_users", len(userList))

	for _, user := range userList {
		for _, userStats := range user.Stats {
			if userStats.Playcount <= 0 {
				continue
			}

			globalRank, err := app.Rankings.GlobalRank(user.Id, userStats.Mode)
			if err != nil {
				return fmt.Errorf("failed to fetch global rank for user %d mode %d: %w", user.Id, userStats.Mode, err)
			}

			peakRank, err := app.Repositories.Histories.FetchPeakGlobalRank(user.Id, userStats.Mode)
			if err != nil {
				return fmt.Errorf("failed to fetch peak rank for user %d mode %d: %w", user.Id, userStats.Mode, err)
			}

			if userStats.Rank != globalRank {
				userStats.Rank = globalRank
				if _, err := app.Repositories.Stats.Update(userStats, "rank"); err != nil {
					return fmt.Errorf("failed to update current rank for user %d mode %d: %w", user.Id, userStats.Mode, err)
				}

				if !app.Config.FrozenRankUpdates {
					if err := app.Repositories.Histories.UpdateRank(userStats, user.Country, app.Rankings); err != nil {
						return fmt.Errorf("failed to update rank history for user %d mode %d: %w", user.Id, userStats.Mode, err)
					}
				}
			}

			if userStats.PeakRank != peakRank {
				userStats.PeakRank = peakRank
				if _, err := app.Repositories.Stats.Update(userStats, "peak_rank"); err != nil {
					return fmt.Errorf("failed to update peak rank for user %d mode %d: %w", user.Id, userStats.Mode, err)
				}
			}
		}

		logger.Info(
			"Updated user",
			"user_id", user.Id, "username", user.Name,
		)
	}
	return nil
}
