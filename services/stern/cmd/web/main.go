package main

import (
	"log/slog"
	"os"

	"github.com/osuTitanic/titanic-go/internal/state"
	"github.com/osuTitanic/titanic-go/services/stern/internal/routes"
	"github.com/osuTitanic/titanic-go/services/stern/internal/server"
	"github.com/osuTitanic/titanic-go/services/stern/internal/templates"
)

func main() {
	app, err := state.NewState(".env")
	if err != nil {
		slog.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}
	defer app.Close()

	engine, err := templates.NewEngine(app.Config)
	if err != nil {
		slog.Error("Failed to initialize templates", "error", err)
		os.Exit(1)
	}

	server := server.NewServer(app.Config.FrontendHost, app.Config.FrontendPort, "stern", app, engine)
	server.Handle("GET /", routes.Home)
	server.Serve()
}
