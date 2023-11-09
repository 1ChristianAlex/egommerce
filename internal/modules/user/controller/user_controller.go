package user_controller

import (
	"net/http"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	user_service "khrix/egommerce/internal/modules/user/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *user_service.UserService
}

func NewModule(router *gin.RouterGroup, userServicer *user_service.UserService) {
	controller := &UserController{
		userService: userServicer,
	}

	router.POST("/create", controller.CreateNewUser)
}

func (controller *UserController) CreateNewUser(context *gin.Context) {
	var userToCreate CreateUserInputDto

	if err := context.ShouldBindJSON(&userToCreate); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*models.User]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	go func() {
		result, err := controller.userService.CreateNewUser(&models.User{
			Username: userToCreate.Username,
			Password: userToCreate.Password,
			Name:     userToCreate.Name,
			Email:    userToCreate.Email,
			Birthday: userToCreate.Birthday,
		})
		if err != nil {
			context.JSON(http.StatusBadRequest, &response.ResponseResult[*models.User]{Result: nil, ErrorMessage: err.Error()})
			return
		}

		context.JSON(http.StatusCreated, &response.ResponseResult[*models.User]{Result: result, ErrorMessage: ""})
	}()
}
