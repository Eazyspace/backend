package controller

import (
	"github.com/labstack/echo/v4"
)

type RootController struct{}

func NewRootController() *RootController {
	return &RootController{}
}

func (c *RootController) InitRouting(g *echo.Group) error {
	g.GET("", c.HealthCheck)
	return nil
}

func (c *RootController) HealthCheck(ctx echo.Context) error {
	ctx.String(200, "Eazyspace backend!")
	return nil
}
