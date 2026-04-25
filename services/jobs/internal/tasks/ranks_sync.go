package tasks

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/osuTitanic/titanic-go/internal/constants"
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

			// Check if rank has desynced from redis -> db & update if necessary
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

			// We want at least 1 rank history update per day
			needsHistoryUpdate := userRequiresHistoryUpdate(
				user.Id,
				userStats.Mode,
				app,
			)
			if needsHistoryUpdate && !app.Config.FrozenRankUpdates {
				if err := app.Repositories.Histories.UpdateRank(userStats, user.Country, app.Rankings); err != nil {
					return fmt.Errorf("failed to update rank history for user %d mode %d: %w", user.Id, userStats.Mode, err)
				}
				logger.Info(
					"Added rank history entry",
					"user_id", user.Id, "mode", userStats.Mode,
				)
			}

			// Keep only the most recent rank history entry per day to save storage space
			err = cleanupRankHistory(
				user.Id,
				userStats.Mode,
				app,
			)
			if err != nil {
				return fmt.Errorf("failed to clean up rank history for user %d mode %d: %w", user.Id, userStats.Mode, err)
			}
		}

		logger.Info(
			"Updated user",
			"user_id", user.Id, "username", user.Name,
		)
	}
	return nil
}

// A user needs to have at least 1 rank history update per day
const RankHistoryUpdateInterval = 24 * 60 * 60

func userRequiresHistoryUpdate(userId int, mode constants.Mode, app *state.State) bool {
	lastUpdate, err := app.Repositories.Histories.FetchLastRankHistoryEntry(userId, mode)
	if err != nil {
		return false
	}
	if lastUpdate == nil {
		return true
	}
	return time.Since(lastUpdate.Time) >= RankHistoryUpdateInterval*time.Second
}

func cleanupRankHistory(userId int, mode constants.Mode, app *state.State) error {
	// We only keep one rank history entry per day
	entries, err := app.Repositories.Histories.FetchRecentRankHistoryEntries(
		userId, mode,
		RankHistoryUpdateInterval*time.Second,
	)
	if err != nil {
		return fmt.Errorf("failed to fetch recent rank history entries for user %d: %w", userId, err)
	}
	if len(entries) <= 1 {
		return nil
	}

	// Delete all but the most recent entry
	for _, entry := range entries[1:] {
		_, err := app.Repositories.Histories.DeleteRankHistoryEntry(userId, mode, entry.Time)
		if err != nil {
			return fmt.Errorf("failed to delete rank history entry %v: %w", entry.Time, err)
		}
	}
	return nil
}
