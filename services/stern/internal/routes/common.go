package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/osuTitanic/titanic-go/internal/state"
	"github.com/osuTitanic/titanic-go/services/stern/internal/server"
	"github.com/osuTitanic/titanic-go/services/stern/internal/templates"
)

func NotFound(ctx *server.Context) {
	ctx.RenderTemplate(
		http.StatusNotFound, "errors/404",
		BuildDefaultView(ctx.State),
	)
}

func InternalServerError(ctx *server.Context) {
	if ctx.Templates == nil {
		ctx.Logger.Error("Failed to render template", "template", "errors/500", "error", "templates engine is not configured")
		templates.InternalServerErrorFallback(ctx.Response)
		return
	}

	body, err := ctx.Templates.Render("errors/500", BuildDefaultView(ctx.State))
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

func BuildDefaultView(state *state.State) templates.DefaultView {
	return templates.DefaultView{
		Stats:  BuildStatistics(state),
		Config: state.Config,
	}
}

func BuildStatistics(state *state.State) (stats templates.Statistics) {
	stats = templates.Statistics{
		TotalUsers:  0,
		OnlineUsers: 0,
		TotalScores: 0,
	}

	values, err := state.Redis.MGet(context.TODO(),
		"bancho:totalusers",
		"bancho:activity:osu",
		"bancho:totalscores",
	).Result()
	if err != nil {
		state.Logger.Error("Failed to fetch statistics from redis", "error", err)
		return stats
	}

	if totalUsers, ok := values[0].(string); ok {
		stats.TotalUsers, _ = strconv.Atoi(totalUsers)
	}
	if onlineUsers, ok := values[1].(string); ok {
		stats.OnlineUsers, _ = strconv.Atoi(onlineUsers)
	}
	if totalScores, ok := values[2].(string); ok {
		stats.TotalScores, _ = strconv.Atoi(totalScores)
	}
	return stats
}
