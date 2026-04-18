package routes

import (
	"net/http"

	"github.com/osuTitanic/titanic-go/internal/config"
	"github.com/osuTitanic/titanic-go/services/stern/internal/server"
	"github.com/osuTitanic/titanic-go/services/stern/internal/templates"
)

func NotFound(ctx *server.Context) {
	ctx.RenderTemplate(
		http.StatusNotFound, "errors/404",
		BuildDefaultView(ctx.State.Config),
	)
}

func BuildDefaultView(cfg *config.Config) templates.DefaultView {
	return templates.DefaultView{
		Stats:  BuildStatistics(),
		Config: cfg,
	}
}

func BuildStatistics() templates.Statistics {
	// TODO: Populate live website stats
	return templates.Statistics{
		TotalUsers:  0,
		OnlineUsers: 0,
		TotalScores: 0,
	}
}
