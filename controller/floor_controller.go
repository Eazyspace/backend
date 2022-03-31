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

var floorErrorPath string = "controller/floor.go: "

type FloorController struct {
	FloorService *service.FloorService
}

func NewFloorController(floorService *service.FloorService) *FloorController {
	return &FloorController{FloorService: floorService}
}

func (c *FloorController) InitRouting(g *echo.Group) error {
	g.GET("", c.GetFloor)
	g.GET("/test", c.Test)
	g.POST("", c.CreateFloor)
	return nil
}

func (c *FloorController) Test(ctx echo.Context) error {
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    false,
		Message: "Floor: OK",
	})
}

func (c *FloorController) CreateFloor(ctx echo.Context) error {

	var input model.Floor

	err := api.GetContent(ctx, &input)

	if err != nil {
		api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", floorErrorPath, err.Error()),
			Data:    false,
		})
	}

	var datas []model.Floor
	// if err != nil {
	// 	api.Respond(ctx, &enum.APIResponse{
	// 		Status:  enum.APIStatus.Error,
	// 		Message: fmt.Sprintf("%s: %s", errorPath, err.Error()),
	// 		Data:    false,
	// 	})
	// }

	// createdRoom, err := c.RoomService.Create(&model.Room{RoomCode: "TEST-001", RoomName: "Room Test", MaxCapacity: 1000})
	createdFloor, err := c.FloorService.Create(&input)
	if err != nil {
		api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", floorErrorPath, err.Error()),
			Data:    false,
		})
	}
	datas = append(datas, *createdFloor)
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Created successfully",
		Data:    datas,
	})

}

// controller handles input & params, (maybe) validate params
func (c *FloorController) GetFloor(ctx echo.Context) error {
	var input model.Floor
	var datas []model.Floor

	param := ctx.QueryParams().Get("q")
	if param == "" {
		param = "{}"
	}

	// convert string -> struct object
	paramErr := json.Unmarshal([]byte(param), &input)

	if paramErr != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("floor_controller/FloorController: paramErr %s", paramErr),
		})
	}

	// return api.Respond(ctx, &enum.APIResponse{
	// 	Status:  enum.APIStatus.Ok,
	// 	Message: "OK",
	// 	Data:    []model.Room{input},
	// })

	datas, err := c.FloorService.Read(&input)

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", floorErrorPath, err.Error()),
		})
	}

	if len(datas) == 0 {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.NotFound,
			Message: fmt.Sprintf("No floor with id %s is found", ctx.QueryParams().Get("id")),
		})
	}

	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "OK",
		Data:    datas,
	})
}
