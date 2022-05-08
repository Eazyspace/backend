package controller

import (
	"encoding/json"
	"fmt"

	"github.com/Eazyspace/api"
	"github.com/Eazyspace/enum"
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/service"
	"github.com/labstack/echo/v4"
)

var adminErrorPath string = "controller/admin.go: "

type AdminController struct {
	AdminService *service.AdminService
}

func NewAdminController(adminService *service.AdminService) *AdminController {
	return &AdminController{AdminService: adminService}
}

func (c *AdminController) InitRouting(g *echo.Group) error {
	g.GET("/test", c.Test)
	g.POST("/accept-request", c.AcceptRequest)
	g.POST("/decline-request", c.DeclineRequest)
	return nil
}

func (c *AdminController) Test(ctx echo.Context) error {
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    false,
		Message: "Admin: OK",
	})
}

func (c *AdminController) AcceptRequest(ctx echo.Context) error {
	body := ctx.Request().Body
	var request *model.Request
	decoder := json.NewDecoder(body)
	decoder.Decode(&request)
	request.Status = 2
	data, err := c.AdminService.UpdateStatus(request)
	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("admin_controller/AdminController: %s", err),
		})
	}
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "OK",
		Data:    data,
	})
}

func (c *AdminController) DeclineRequest(ctx echo.Context) error {
	body := ctx.Request().Body
	var request *model.Request
	decoder := json.NewDecoder(body)
	decoder.Decode(&request)
	request.Status = 3
	data, err := c.AdminService.UpdateStatus(request)
	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintf("admin_controller/AdminController: err %s", err),
		})
	}
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "OK",
		Data:    data,
	})
}
