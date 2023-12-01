package di

import (
	"khrix/egommerce/internal/modules/user/dto"
)

type UserService interface {
	CreateNewUser(userModel *dto.UserInputDto) (newUser *dto.UserOutputDto, err error)
	TryLogin(email, password string) (newUser *dto.UserOutputDto, err error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type JwtService interface {
	NewClains(user dto.UserOutputDto) (string, error)
	FromClains(tokenString string) (*dto.UserOutputTokenDto, error)
	IsValid(tokenString string) error
}
