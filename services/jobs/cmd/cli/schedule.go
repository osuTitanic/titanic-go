package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/osuTitanic/titanic-go/internal/state"
	"github.com/osuTitanic/titanic-go/services/jobs/internal/scheduler"
)

func ScheduleTask(app *state.State, s *scheduler.Scheduler, name string, interval int, intervalAt string) error {
	taskFunc, ok := availableTasks[name]
	if !ok {
		return fmt.Errorf("unknown task: %s", name)
	}

	period := time.Duration(interval) * time.Second
	schedule := scheduler.Every(period)

	task := s.Add(schedule, taskFunc)
	task.SetLogger(name)

	if intervalAt != "" {
		task.Schedule = schedule.At(intervalAt)
	}

	app.Logger.Info("Scheduled task", "name", name, "interval", interval)
	return nil
}

func StartSchedulerAndWait(app *state.State, s *scheduler.Scheduler) {
	s.Start(app)
	app.Logger.Info("Scheduler started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	app.Logger.Info("Shutting down...")
	s.Stop()
}

func RunSchedulerFile(app *state.State, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read scheduler file: %w", err)
	}

	var configs []SchedulerConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return fmt.Errorf("parse scheduler file: %w", err)
	}
	s := scheduler.New()

	for _, config := range configs {
		if config.Interval <= 0 {
			app.Logger.Error("Interval must be > 0 for scheduled tasks", "name", config.Name)
			continue
		}

		intervalAt := ""
		if config.IntervalAt != nil {
			intervalAt = *config.IntervalAt
		}

		if err := ScheduleTask(app, s, config.Name, config.Interval, intervalAt); err != nil {
			app.Logger.Error("Failed to schedule task", "name", config.Name, "error", err)
		}
	}

	StartSchedulerAndWait(app, s)
	return nil
}
