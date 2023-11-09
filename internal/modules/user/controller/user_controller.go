package controller

import (
	"net/http"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/modules/user/repository/entities"
	"khrix/egommerce/internal/modules/user/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewModule(router *gin.RouterGroup, userServicer *service.UserService) {
	controller := &UserController{
		userService: userServicer,
	}

	router.POST("/create", controller.CreateNewUser)
}

func (controller *UserController) CreateNewUser(context *gin.Context) {
	var userToCreate CreateUserInputDto

	if err := context.ShouldBindJSON(&userToCreate); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*entities.User]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	go func() {
		result, err := controller.userService.CreateNewUser(&entities.User{
			Username: userToCreate.Username,
			Password: userToCreate.Password,
			Name:     userToCreate.Name,
			Email:    userToCreate.Email,
			Birthday: userToCreate.Birthday,
		})
		if err != nil {
			context.JSON(http.StatusBadRequest, &response.ResponseResult[*entities.User]{Result: nil, ErrorMessage: err.Error()})
			return
		}

		context.JSON(http.StatusCreated, &response.ResponseResult[*entities.User]{Result: result, ErrorMessage: ""})
	}()
}
