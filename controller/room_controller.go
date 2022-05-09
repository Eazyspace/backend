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

var roomErrorPath string = "controller/room.go: "

type RoomController struct {
	RoomService *service.RoomService
}

func NewRoomController(roomService *service.RoomService) *RoomController {
	return &RoomController{RoomService: roomService}
}

func (c *RoomController) InitRouting(g *echo.Group) error {
	g.GET("", c.GetRoom)
	g.GET("/test", c.Test)
	g.POST("/create", c.CreateRoom)
	g.PUT("/update", c.UpdateRoom)
	g.POST("/book", c.BookRoom)
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
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", roomErrorPath, err.Error()),
			Data:    false,
		})
	}

	if input.FloorID == 0 ||
		input.RoomName == "" ||
		input.RoomLength == 0 ||
		input.RoomWidth == 0 ||
		input.MaxCapacity == 0 {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Forbidden,
			Message: "Missing param (floorId, roomName, roomLength, roomWidth, maxCapacity)",
			Data:    false,
		})
	}

	var datas []model.Room
	createdRoom, err := c.RoomService.Create(&input)
	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", roomErrorPath, err.Error()),
			Data:    false,
		})
	}
	datas = append(datas, *createdRoom)
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Create room successfully",
		Data:    datas,
	})
}

func (c *RoomController) UpdateRoom(ctx echo.Context) error {
	var input model.Room
	err := api.GetContent(ctx, &input)
	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", roomErrorPath, err.Error()),
			Data:    false,
		})
	}

	if input.RoomID == 0 {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: "Missing roomId",
			Data:    false,
		})
	}

	var datas []model.Room
	updatedRoom, err := c.RoomService.Update(&input)
	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", roomErrorPath, err.Error()),
			Data:    false,
		})
	}
	datas = append(datas, *updatedRoom)
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Update room successfully",
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
			Message: fmt.Sprintf("%s: %s", roomErrorPath, err.Error()),
		})
	}

	if len(datas) == 0 {
		return api.Respond(ctx, &enum.APIResponse{
			Status: enum.APIStatus.NotFound,
			Message: fmt.Sprintf("No room with id %s or floorId %s is found",
				ctx.QueryParams().Get("id"), ctx.QueryParams().Get("floorId")),
		})
	}

	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "OK",
		Data:    datas,
	})
}

func (c *RoomController) BookRoom(ctx echo.Context) error {
	var input model.Request

	err := api.GetContent(ctx, &input)

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", roomErrorPath, err.Error()),
			Data:    false,
		})
	}

	if input.RoomID == 0 ||
		input.UserID == 0 ||
		input.StartTime.IsZero() ||
		input.EndTime.IsZero() ||
		input.Description == "" ||
		input.EventName == "" {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Forbidden,
			Message: "Missing param (roomId, userId, startTime, endTime, description, eventName)",
			Data:    false,
		})
	}

	var datas []model.Request

	// return api.Respond(ctx, &enum.APIResponse{
	// 	Status:  enum.APIStatus.Ok,
	// 	Message: "OK",
	// 	Data:    []model.Request{input},
	// })

	createdRequest, errString := c.RoomService.Book(&input)
	if errString != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", roomErrorPath, errString),
			Data:    false,
		})
	}

	datas = append(datas, *createdRequest)
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Booking request successfully",
		Data:    datas,
	})
}
