package echo_server

import "github.com/labstack/echo/v4"

type Config struct {
	server *echo.Echo
	router *Router
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Handler(handler Handler) *Config {
	c.server = echo.New()
	c.router = NewRouter(c.server, handler)
	return c
}

func (c *Config) Run() {
	e := c.server
	c.router.Exec()
	e.Logger.Fatal(e.Start(":1324"))
}
