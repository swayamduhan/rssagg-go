package router

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func TestRoutes(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from test routes")
}