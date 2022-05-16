package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Eazyspace/api"
	"github.com/Eazyspace/enum"
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/service"
	"github.com/golang-jwt/jwt"
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
	g.POST("/check-in", c.CheckIn)
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
	var datas []model.Request

	param := ctx.QueryParams().Get("q")
	if param == "" {
		param = "{}"
	}

	input := make(map[string]interface{})
	errMap := json.Unmarshal([]byte(param), &input)
	if errMap != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("request_controller/RequestController: paramErr %s", errMap),
		})
	}

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

func (c *RequestController) CheckIn(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if len(token) == 0 {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Unauthorized,
			Message: "A token is required",
		})
	}
	if strings.Contains(token, "Bearer ") {
		token = strings.SplitAfter(token, "Bearer ")[1]
	}

	jwtToken, err := jwt.Parse(token,
		func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Unauthorized,
			Data:    err,
			Message: "Invalid token",
		})
	}

	var tokenObject model.Token
	stringified, _ := json.Marshal(jwtToken.Claims)

	json.Unmarshal([]byte(stringified), &tokenObject)

	var input model.Request

	err = api.GetContent(ctx, &input)

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("request_controller/RequestController: paramErr %s", err),
		})
	}

	result, err := c.RequestService.CheckIn(&input, tokenObject.UserID)

	if err != nil {
		if err.Error() == "unauthorized" {
			return api.Respond(ctx, &enum.APIResponse{
				Status:  enum.APIStatus.Unauthorized,
				Message: "Please use the account that made this request to checkin",
			})
		}
		if err.Error() == "not found" {
			return api.Respond(ctx, &enum.APIResponse{
				Status:  enum.APIStatus.Unauthorized,
				Message: "Request not found",
			})
		}
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("request_controller/RequestController: paramErr %s", err),
		})
	}
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: strconv.FormatBool(result),
	})
}
