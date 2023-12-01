package di

import (
	"khrix/egommerce/internal/modules/user/dto"
	"khrix/egommerce/internal/modules/user/repository/entities"
)

type UserMapper interface {
	ToDto(item entities.User) dto.UserOutputDto
	ToEntity(item dto.UserInputDto) entities.User
}
