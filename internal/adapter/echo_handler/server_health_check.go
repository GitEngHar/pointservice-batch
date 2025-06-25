package echo_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pointservice-batch/internal/adapter/client_handler"
	"pointservice-batch/internal/infra/client"
)

type ServerHealthCheckHandler struct{}

func NewServerHealthCheckHandler() *ServerHealthCheckHandler {
	return &ServerHealthCheckHandler{}
}

func (h *ServerHealthCheckHandler) Handle(ctx echo.Context) error {
	config := client.NewConfig()
	handler := client_handler.NewHttpHealthCheckHandler()
	response, err := client.NewHttpClient(config, handler).Client().Start()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, []string{})
	}
	return ctx.JSON(http.StatusOK, []string{response})
}
