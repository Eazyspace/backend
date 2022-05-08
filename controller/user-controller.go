package controller

import (
	"encoding/json"
	"fmt"

	"github.com/Eazyspace/api"
	"github.com/Eazyspace/enum"
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var userErrorPath string = "controller/user.go: "

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) InitRouting(g *echo.Group) error {
	g.GET("", c.GetUser, api.CheckAuthorization())
	g.GET("/test", c.Test)
	g.POST("/login", c.Login)
	g.POST("/sign-up", c.SignUp)
	return nil
}

func (c *UserController) Test(ctx echo.Context) error {
	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    false,
		Message: "User: OK",
	})
}

func (c *UserController) GetUser(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Token)

	var input model.User
	input.UserID = claims.UserID

	var datas []model.User
	datas, errString := c.UserService.Read(&input)

	if errString != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", userErrorPath, errString),
			Data:    false,
		})
	}

	result, err := json.Marshal(datas[0])

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", userErrorPath, err.Error()),
			Data:    false,
		})
	}
	m := make(map[string]interface{})
	json.Unmarshal(result, &m)
	delete(m, "password")

	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    m,
		Message: "User: OK",
	})
}

func (c *UserController) Login(ctx echo.Context) error {
	var input model.User

	err := api.GetContent(ctx, &input)

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", userErrorPath, err.Error()),
			Data:    false,
		})
	}

	if input.AcademicID == "" ||
		input.Password == "" {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Unauthorized,
			Message: "Missing username or password",
			Data:    false,
		})
	}

	tokenString, errString := c.UserService.Login(&input)
	if errString != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", userErrorPath, errString),
			Data:    false,
		})
	}

	return api.Respond(ctx, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Message: "Log in successfully",
		Data: echo.Map{
			"token": tokenString,
		},
	})
}

func (c *UserController) SignUp(ctx echo.Context) error {
	var user *model.User

	err := api.GetContent(ctx, &user)

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", userErrorPath, err.Error()),
			Data:    false,
		})
	}

	data, err := c.UserService.Create(user)
	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Invalid,
			Message: fmt.Sprintf("%s: %s", userErrorPath, err),
		})
	}

	result, err := json.Marshal(data)

	if err != nil {
		return api.Respond(ctx, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("%s: %s", userErrorPath, err.Error()),
			Data:    false,
		})
	}
	m := make(map[string]interface{})
	json.Unmarshal(result, &m)
	delete(m, "password")

	return api.Respond(ctx, &enum.APIResponse{
		Status: enum.APIStatus.Ok,
		Data:   m,
	})
}
