package service

import (
	"khrix/egommerce/internal/modules/user/repository"
	"khrix/egommerce/internal/modules/user/repository/entities"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (service UserService) CreateNewUser(userModel *entities.User) (newUser *entities.User, error error) {
	_, err := service.repository.CreateUser(userModel)

	return userModel, err
}
