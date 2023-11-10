package controller

import (
	"net/http"

	coreDi "khrix/egommerce/internal/core/di"
	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	"khrix/egommerce/internal/modules/user/di"
	"khrix/egommerce/internal/modules/user/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService di.UserService
	jwtService  di.JwtService
	authHelper  coreDi.AuthHelper
}

func NewModule(router *gin.RouterGroup, userServicer di.UserService, jwtService di.JwtService, authHelper coreDi.AuthHelper) {
	controller := &UserController{
		userService: userServicer,
		jwtService:  jwtService,
		authHelper:  authHelper,
	}

	apiRouter := router.Group("api", controller.authHelper.JwtMiddleware)

	router.POST("/login", controller.DoLogin)

	apiRouter.POST("/create", controller.CreateNewUser)
}

func (controller *UserController) CreateNewUser(context *gin.Context) {
	var userToCreate dto.CreateUserInputDto

	if err := context.ShouldBindJSON(&userToCreate); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.UserOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	result, err := controller.userService.CreateNewUser(&userToCreate)
	if err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.UserOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, &response.ResponseResult[*dto.UserOutputDto]{Result: result, ErrorMessage: ""})
}

func (controller *UserController) DoLogin(context *gin.Context) {
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
	},
	)
}
