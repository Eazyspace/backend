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

var requestErrorPath string = "controller/request.go: "

type RequestController struct {
	RequestService *service.RequestService
}

func NewRequestController(requestService *service.RequestService) *RequestController {
	return &RequestController{RequestService: requestService}
}

func (c *RequestController) InitRouting(g *echo.Group) error {
	g.GET("", c.GetRequest)
	g.GET("/test", c.Test)
	g.POST("", c.CreateRequest)
	return nil
}

func (c *RequestController) Test(ctx echo.Context) error {
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    false,
		Message: "Request: OK",
	})
}

func (c *RequestController) CreateRequest(ctx echo.Context) error {

	var input model.Request

	err := api.GetContent(ctx, &input)

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", requestErrorPath, err.Error()),
			Data:    false,
		})
	}

	var datas []model.Request
	// if err != nil {
	// 	api.Respond(ctx, &enum.APIResponse{
	// 		Status:  enum.APIStatus.Error,
	// 		Message: fmt.Sprintf("%s: %s", errorPath, err.Error()),
	// 		Data:    false,
	// 	})
	// }

	// createdRequest, err := c.RequestService.Create(&model.Request{RequestCode: "TEST-001", RequestName: "Request Test", MaxCapacity: 1000})
	createdRequest, err := c.RequestService.Create(&input)
	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", requestErrorPath, err.Error()),
			Data:    false,
		})
	}
	datas = append(datas, *createdRequest)
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Created successfully",
		Data:    datas,
	})

}

// controller handles input & params, (maybe) validate params
func (c *RequestController) GetRequest(ctx echo.Context) error {
	var input model.Request
	var datas []model.Request

	param := ctx.QueryParams().Get("q")
	if param == "" {
		param = "{}"
	}

	paramErr := json.Unmarshal([]byte(param), &input)

	if paramErr != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("request_controller/RequestController: paramErr %s", paramErr),
		})
	}

	// return api.Respond(ctx, &enum.APIResponse{
	// 	Status:  enum.APIStatus.Ok,
	// 	Message: "OK",
	// 	Data:    []model.Request{input},
	// })

	datas, err := c.RequestService.Read(&input)

	if err != nil {
		api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", requestErrorPath, err.Error()),
		})
	}

	if len(datas) == 0 {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.NotFound,
			Message: fmt.Sprintf("No request found"),
		})
	}

	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "OK",
		Data:    datas,
	})
}
