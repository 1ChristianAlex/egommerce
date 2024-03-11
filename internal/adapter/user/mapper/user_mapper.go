package mapper

import (
	"khrix/egommerce/internal/application/user/dto"
	"khrix/egommerce/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m UserMapper) ToDto(item models.User) *dto.UserOutputDto {
	return &dto.UserOutputDto{
		UserID:   int32(item.ID),
		Username: item.Username,
		Name:     item.Name,
		Email:    item.Email,
		Birthday: item.Birthday,
	}
}

func (m UserMapper) ToEntity(item dto.UserInputDto) *models.User {
	return &models.User{
		Username: item.Username,
		Password: item.Password,
		Name:     item.Name,
		Email:    item.Email,
		Birthday: item.Birthday,
		Model:    gorm.Model{ID: item.ID},
	}
}
