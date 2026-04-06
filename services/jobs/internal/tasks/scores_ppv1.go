package tasks

import (
	"log/slog"
	"sort"
	"sync"

	"github.com/osuTitanic/titanic-go/internal/schemas"
	"github.com/osuTitanic/titanic-go/internal/state"
)

// UpdatePPv1 updates ppv1 calculations for all users
func UpdatePPv1(app *state.State, logger *slog.Logger) error {
	if app.Config.FrozenPPv1Updates {
		logger.Info("ppv1 updates are disabled, skipping...")
		return nil
	}

	criteria := map[string]any{
		"restricted = ?": false,
		"activated = ?":  true,
	}
	userList, err := app.Repositories.Users.Many(criteria, "Stats")
	if err != nil {
		return err
	}
	logger.Info("Updating ppv1 calculations...", "total_users", len(userList))

	sort.Slice(userList, func(i, j int) bool {
		return resolveUserPPv1(userList[i]) > resolveUserPPv1(userList[j])
	})

	var wg sync.WaitGroup

	performUpdate := func(u *schemas.User) {
		defer wg.Done()
		err := updatePPv1ForUser(app, logger, u)
		if err != nil {
			logger.Error("Failed to update user", "id", u.Id, "error", err)
		}
	}
	for _, user := range userList {
		wg.Add(1)
		go performUpdate(user)
	}

	wg.Wait()
	return nil
}

func updatePPv1ForUser(app *state.State, logger *slog.Logger, user *schemas.User) error {
	for _, stats := range user.Stats {
		if stats.Playcount <= 0 {
			continue
		}

		bestScores, err := app.Repositories.Scores.FetchBest(
			user.Id,
			stats.Mode,
			!app.Config.ApprovedMapRewards,
			"Beatmap",
		)
		if err != nil {
			return err
		}
		if len(bestScores) == 0 {
			continue
		}

		stats.PPv1, err = app.PPv1.RecalculateWeightFromScores(bestScores)
		if err != nil {
			return err
		}

		app.Repositories.Stats.Update(stats, "ppv1")
		app.Rankings.Update(stats, user.Country)
		// TODO: Update rank history

		logger.Debug(
			"ppv1 update",
			"id", user.Id, "name", user.Name,
			"mode", stats.Mode, "ppv1", stats.PPv1,
		)
	}

	logger.Info("Updated ppv1 for user", "name", user.Name, "id", user.Id)
	return nil
}

func resolveUserPPv1(user *schemas.User) float64 {
	var totalPPv1 float64
	for _, stats := range user.Stats {
		totalPPv1 += stats.PPv1
	}
	return totalPPv1
}
