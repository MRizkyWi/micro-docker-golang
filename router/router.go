package router

import (
	"micro-echo/api"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	api.MainGroup(e)

	return e
}
