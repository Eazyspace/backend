package controller

import "github.com/labstack/echo/v4"

type BaseController interface {
	InitRouting(g *echo.Group) error
}
