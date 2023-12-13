package users

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	CreateHandler(c echo.Context) error
	DetailHandler(c echo.Context) error
}
