package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/osuTitanic/titanic-go/internal/state"
)

// Server is the main struct that holds the state for an http server.
type Server struct {
	Host   string
	Port   int
	Name   string
	State  *state.State
	Router *mux.Router
	Logger *slog.Logger
}

func NewServer(host string, port int, name string, state *state.State) *Server {
	return &Server{
		Host:   host,
		Port:   port,
		Name:   name,
		State:  state,
		Logger: slog.Default().With("component", name),
		Router: mux.NewRouter(),
	}
}

// Context is a struct that holds the request context for each endpoint call.
type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	State    *state.State
	Vars     map[string]string
}

func (ctx *Context) IP() string {
	return GetRequestIP(ctx.Request)
}

// Serve starts the HTTP server and listens for incoming requests.
func (server *Server) Serve() {
	bind := fmt.Sprintf(
		"%s:%d",
		server.Host,
		server.Port,
	)
	server.Logger.Info(
		"Listening for requests",
		"host", server.Host,
		"port", server.Port,
	)

	err := http.ListenAndServe(bind, server.LoggingMiddleware(server.Router))
	if err != nil {
		server.Logger.Error("Failed to start server", "error", err)
		return
	}
}

// ResponseContext is a wrapper around http.ResponseWriter that
// allows us to capture the status code of a response.
type ResponseContext struct {
	w http.ResponseWriter
	s int
}

func (rc *ResponseContext) Header() http.Header {
	return rc.w.Header()
}

func (rc *ResponseContext) Write(b []byte) (int, error) {
	return rc.w.Write(b)
}

func (rc *ResponseContext) WriteHeader(status int) {
	rc.s = status
	rc.w.WriteHeader(status)
}

func (rc *ResponseContext) Status() int {
	if rc.s == 0 {
		return http.StatusOK
	}
	return rc.s
}

// ContextMiddleware creates a new Context struct for each request.
func (server *Server) ContextMiddleware(handler func(*Context)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := &Context{
			Response: w,
			Request:  r,
			State:    server.State,
			Vars:     mux.Vars(r),
		}

		w.Header().Set("Server", server.Name)
		handler(context)
	}
}

// LoggingMiddleware logs the details of each request.
func (server *Server) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rc := &ResponseContext{w: w}
		start := time.Now()
		next.ServeHTTP(rc, r)
		time := time.Since(start)

		server.Logger.Info(
			fmt.Sprintf("%s %s", r.Method, r.RequestURI),
			"ip", GetRequestIP(r),
			"status", rc.Status(),
			"duration", time.String(),
		)
	})
}
