package http

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-unit-test/internal/users"
)

func MapRoutes(e *echo.Echo, h users.Handlers) {
	e.POST("/users", h.CreateHandler)
	e.GET("/users/:id", h.DetailHandler)
}
