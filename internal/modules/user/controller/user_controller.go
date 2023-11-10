package controller

import (
	"net/http"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/modules/user/di"
	"khrix/egommerce/internal/modules/user/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService di.UserService
}

func NewUserModule(router *gin.RouterGroup, userServicer di.UserService) {
	controller := &UserController{
		userService: userServicer,
	}

	router.POST("/create", controller.CreateNewUser)
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
