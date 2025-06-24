package echo_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Handle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, []string{"ok"})
}
