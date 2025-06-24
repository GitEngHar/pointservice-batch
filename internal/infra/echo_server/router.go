package echo_server

import "github.com/labstack/echo/v4"

type Handler interface {
	Handle(ctx echo.Context) error
}

type Router struct {
	server  *echo.Echo
	handler Handler
}

func NewRouter(server *echo.Echo, handler Handler) *Router {
	return &Router{
		server:  server,
		handler: handler,
	}
}

func (r *Router) Exec() {
	e := r.server
	e.GET("/", r.handler.Handle)
}
