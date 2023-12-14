package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-unit-test/internal/users"
	"github.com/martinyonatann/go-unit-test/internal/users/dtos"
)

type handlers struct {
	uc users.UseCases
}

func NewHandlers(uc users.UseCases) users.Handlers {
	return &handlers{uc: uc}
}

func (h *handlers) CreateHandler(c echo.Context) error {
	var request dtos.Users
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = h.uc.CreateUser(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}
func (h *handlers) DetailHandler(c echo.Context) error {
	var request dtos.UserDetailRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	detail, err := h.uc.GetUserDetail(c.Request().Context(), request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, detail)
}
