package user

import (
	"khrix/egommerce/internal/application/user/dto"
	"khrix/egommerce/internal/port/libs"
)

type UserMapper[R any] interface {
	libs.BaseMapper[dto.UserInputDto, dto.UserOutputDto, R]
}
