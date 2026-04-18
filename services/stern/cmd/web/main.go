package main

import (
	"log/slog"
	"os"

	"github.com/osuTitanic/titanic-go/internal/state"
	"github.com/osuTitanic/titanic-go/services/stern/internal/server"
)

func main() {
	app, err := state.NewState(".env")
	if err != nil {
		slog.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}
	defer app.Close()

	server := server.NewServer(app.Config.FrontendHost, app.Config.FrontendPort, "stern", app)
	server.Serve()
}
