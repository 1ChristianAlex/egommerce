package controller

import (
	"net/http"

	"khrix/egommerce/internal/adapter/libs/base_controller"
	"khrix/egommerce/internal/application/libs/channels"
	user_dto "khrix/egommerce/internal/application/user/dto"
	"khrix/egommerce/internal/port/auth"
	"khrix/egommerce/internal/port/user"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService user.UserService
	jwtService  auth.JwtService
}

func NewAuthController(userServicer user.UserService, jwtService auth.JwtService) *AuthController {
	return &AuthController{
		userService: userServicer,
		jwtService:  jwtService,
	}
}

func (controller *AuthController) DoLogin(context *gin.Context) {
	var loginInput user_dto.LoginInputDto

	if err := context.ShouldBindJSON(&loginInput); err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[*user_dto.UserOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan channels.Resolve[*user_dto.UserOutputDto])
	defer close(channel)

	go func() {
		result, err := controller.userService.TryLogin(
			loginInput.Email,
			loginInput.Password,
		)

		channel <- channels.Resolve[*user_dto.UserOutputDto]{
			Result: result,
			Err:    err,
		}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[*user_dto.UserOutputDto]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	tokenChannel := make(chan channels.Resolve[string])
	defer close(tokenChannel)

	go func() {
		token, errToken := controller.jwtService.NewClains(*resolve.Result)

		tokenChannel <- channels.Resolve[string]{Result: token, Err: errToken}
	}()

	tokenResult := <-tokenChannel

	if tokenResult.Err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[*user_dto.UserOutputDto]{Result: nil, ErrorMessage: tokenResult.Err.Error()})
		return
	}

	context.JSON(http.StatusOK, &base_controller.ResponseResult[interface{}]{
		Result: struct {
			User  *user_dto.UserOutputDto `json:"user" `
			Token string                  `json:"token" `
		}{
			User: resolve.Result, Token: tokenResult.Result,
		},
	})
}
