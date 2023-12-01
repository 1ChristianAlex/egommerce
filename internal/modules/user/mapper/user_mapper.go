package mapper

import (
	"khrix/egommerce/internal/modules/user/dto"
	"khrix/egommerce/internal/modules/user/repository/entities"

	"gorm.io/gorm"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m UserMapper) ToDto(item entities.User) dto.UserOutputDto {
	return dto.UserOutputDto{
		ID:       item.ID,
		Username: item.Username,
		Name:     item.Name,
		Email:    item.Email,
		Birthday: item.Birthday,
	}
}

func (m UserMapper) ToEntity(item dto.UserInputDto) entities.User {
	return entities.User{
		Username: item.Username,
		Password: item.Password,
		Name:     item.Name,
		Email:    item.Email,
		Birthday: item.Birthday,
		Model:    gorm.Model{ID: item.ID},
	}
}
