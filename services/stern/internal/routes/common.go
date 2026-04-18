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

func InternalServerError(ctx *server.Context) {
	if ctx.Templates == nil {
		ctx.Logger.Error("Failed to render template", "template", "errors/500", "error", "templates engine is not configured")
		templates.InternalServerErrorFallback(ctx.Response)
		return
	}

	body, err := ctx.Templates.Render("errors/500", BuildDefaultView(ctx.State.Config))
	if err != nil {
		ctx.Logger.Error("Failed to render template", "template", "errors/500", "error", err)
		templates.InternalServerErrorFallback(ctx.Response)
		return
	}

	ctx.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	ctx.Response.WriteHeader(http.StatusInternalServerError)
	if _, err := ctx.Response.Write(body); err != nil {
		ctx.Logger.Error("Failed to write response body", "template", "errors/500", "error", err)
	}
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
