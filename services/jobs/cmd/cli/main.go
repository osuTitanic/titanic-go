package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/osuTitanic/titanic-go/internal/state"
	"github.com/osuTitanic/titanic-go/services/jobs/internal/tasks"
)

type TaskList map[string]func(*state.State, *slog.Logger) error

func (t *TaskList) List() {
	slog.Info("Available tasks:")
	for name := range availableTasks {
		slog.Info(fmt.Sprintf(" - %s", name))
	}
}

var availableTasks = TaskList{
	"stats_website": tasks.UpdateWebsiteStats,
}

type SchedulerConfig struct {
	Name       string  `json:"name"`
	Interval   int     `json:"interval"` // in seconds
	IntervalAt *string `json:"interval_at"`
}

func main() {
	listFlag := flag.Bool("list", false, "list all tasks")
	nameFlag := flag.String("name", "", "run a specific task by name")
	intervalFlag := flag.Int("interval", 0, "the interval to run the task in (seconds)")
	intervalAtFlag := flag.String("interval-at", "", "specify the time period at which the task should run (e.g. 15:00)")
	fileFlag := flag.String("file", "", "specify the scheduler file")
	flag.Parse()

	if *listFlag {
		availableTasks.List()
		os.Exit(0)
	}

	app, err := state.NewState(".env")
	if err != nil {
		slog.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}
	defer app.Close()

	if *fileFlag != "" {
		if err := RunSchedulerFile(app, *fileFlag); err != nil {
			app.Logger.Error("Failed to run tasks from file", "error", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if *nameFlag != "" {
		err := RunSingleTask(app, *nameFlag, *intervalFlag, *intervalAtFlag)
		if err != nil {
			app.Logger.Error("Failed to run task", "error", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	flag.Usage()
	os.Exit(1)
}
