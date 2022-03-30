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

var errorPath string = "controller/room.go: "

type RoomController struct {
	RoomService *service.RoomService
}

func NewRoomController(roomService *service.RoomService) *RoomController {
	return &RoomController{RoomService: roomService}
}

func (c *RoomController) InitRouting(g *echo.Group) error {
	g.GET("", c.GetRoom)
	g.GET("/test", c.Test)
	g.POST("", c.CreateRoom)
	return nil
}

func (c *RoomController) Test(ctx echo.Context) error {
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    false,
		Message: "Room: OK",
	})
}

func (c *RoomController) CreateRoom(ctx echo.Context) error {

	var input model.Room

	err := api.GetContent(ctx, &input)

	if err != nil {
		api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", errorPath, err.Error()),
			Data:    false,
		})
	}

	var datas []model.Room
	// if err != nil {
	// 	api.Respond(ctx, &enum.APIResponse{
	// 		Status:  enum.APIStatus.Error,
	// 		Message: fmt.Sprintf("%s: %s", errorPath, err.Error()),
	// 		Data:    false,
	// 	})
	// }

	// createdRoom, err := c.RoomService.Create(&model.Room{RoomCode: "TEST-001", RoomName: "Room Test", MaxCapacity: 1000})
	createdRoom, err := c.RoomService.Create(&input)

	datas = append(datas, *createdRoom)
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Created successfully",
		Data:    datas,
	})

}

// controller handles input & params, (maybe) validate params
func (c *RoomController) GetRoom(ctx echo.Context) error {
	var input model.Room
	var datas []model.Room

	param := ctx.QueryParams().Get("q")
	if param == "" {
		param = "{}"
	}
	// convert string -> struct object
	paramErr := json.Unmarshal([]byte(param), &input)

	if paramErr != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("room_controller/RoomController: paramErr %s", paramErr),
		})
	}

	// return api.Respond(ctx, &enum.APIResponse{
	// 	Status:  enum.APIStatus.Ok,
	// 	Message: "OK",
	// 	Data:    []model.Room{input},
	// })

	datas, err := c.RoomService.Read(&input)

	if err != nil {
		api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", errorPath, err.Error()),
		})
	}

	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "OK",
		Data:    datas,
	})
}
