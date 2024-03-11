package user

import "khrix/egommerce/internal/application/user/dto"

type UserService interface {
	CreateNewUser(userModel *dto.UserInputDto) (newUser *dto.UserOutputDto, err error)
	TryLogin(email, password string) (newUser *dto.UserOutputDto, err error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
