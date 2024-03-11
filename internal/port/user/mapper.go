package user

import (
	"khrix/egommerce/internal/application/user/dto"
	"khrix/egommerce/internal/infrastructure/database/models"
)

type UserMapper interface {
	ToDto(item models.User) dto.UserOutputDto
	ToEntity(item dto.UserInputDto) models.User
}
