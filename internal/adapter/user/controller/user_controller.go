package controller

import (
	"net/http"

	base_controller "khrix/egommerce/internal/adapter/libs/base_controller"
	"khrix/egommerce/internal/application/user/dto"
	"khrix/egommerce/internal/port/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user.UserService
}

func NewUserController(userServicer user.UserService) *UserController {
	return &UserController{
		userService: userServicer,
	}
}

func (controller *UserController) CreateNewUser(context *gin.Context) {
	var userToCreate dto.UserInputDto

	if err := context.ShouldBindJSON(&userToCreate); err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[*dto.UserOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	result, err := controller.userService.CreateNewUser(&userToCreate)
	if err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[*dto.UserOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, &base_controller.ResponseResult[*dto.UserOutputDto]{Result: result, ErrorMessage: ""})
}
