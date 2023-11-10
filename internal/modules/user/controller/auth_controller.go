package controller

import (
	"net/http"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	"khrix/egommerce/internal/modules/user/di"
	"khrix/egommerce/internal/modules/user/dto"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService di.UserService
	jwtService  di.JwtService
}

func NewAuthModule(router *gin.RouterGroup, userServicer di.UserService, jwtService di.JwtService) {
	controller := &AuthController{
		userService: userServicer,
		jwtService:  jwtService,
	}

	router.POST("/login", controller.DoLogin)
}

func (controller *AuthController) DoLogin(context *gin.Context) {
	var loginInput dto.LoginInputDto

	if err := context.ShouldBindJSON(&loginInput); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.UserOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[*dto.UserOutputDto])
	defer close(channel)

	go func() {
		result, err := controller.userService.TryLogin(
			loginInput.Email,
			loginInput.Password,
		)

		channel <- models.Resolve[*dto.UserOutputDto]{
			Result: result,
			Err:    err,
		}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.UserOutputDto]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	token, errToken := controller.jwtService.NewClains(*resolve.Result)

	if errToken != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.UserOutputDto]{Result: nil, ErrorMessage: errToken.Error()})
		return
	}

	context.JSON(http.StatusOK, &response.ResponseResult[interface{}]{
		Result: struct {
			User  *dto.UserOutputDto `json:"user" `
			Token string             `json:"token" `
		}{
			User: resolve.Result, Token: token,
		},
	})
}
