package user_service

import (
	"khrix/egommerce/internal/models"
	user_repository "khrix/egommerce/internal/modules/user/repository"
)

type UserService struct {
	repository *user_repository.UserRepository
}

func NewUserService(repository *user_repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (service UserService) CreateNewUser(userModel *models.User) (newUser *models.User, error error) {
	_, err := service.repository.CreateUser(userModel)

	return userModel, err
}
