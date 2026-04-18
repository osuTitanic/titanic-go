package routes

import "github.com/osuTitanic/titanic-go/services/stern/internal/server"

func Home(ctx *server.Context) {
	ctx.Response.WriteHeader(200)
	ctx.Response.Write([]byte("yeah"))
}
