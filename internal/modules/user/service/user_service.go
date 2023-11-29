package service

import (
	"errors"

	"khrix/egommerce/internal/modules/user/di"
	"khrix/egommerce/internal/modules/user/dto"
	"khrix/egommerce/internal/modules/user/repository/entities"
)

type UserService struct {
	repository      di.UserRepository
	passwordService di.PasswordService
}

func NewUserService(repository di.UserRepository, passwordService di.PasswordService) *UserService {
	return &UserService{
		repository:      repository,
		passwordService: passwordService,
	}
}

func (service UserService) CreateNewUser(userInput *dto.CreateUserInputDto) (newUser *dto.UserOutputDto, error error) {
	userExist, _ := service.repository.FindByEmail(userInput.Email)

	if userExist != nil {
		return nil, errors.New("email already exist")
	}

	hash, hashError := service.passwordService.HashPassword(userInput.Password)

	if hashError != nil {
		return nil, errors.New("fail on create user password")
	}

	userModel := entities.User{
		Username: userInput.Username,
		Password: hash,
		Name:     userInput.Name,
		Email:    userInput.Email,
		Birthday: userInput.Birthday,
	}
	user, err := service.repository.CreateUser(&userModel)

	return &dto.UserOutputDto{
		Id:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, err
}

func (service UserService) TryLogin(email, password string) (newUser *dto.UserOutputDto, error error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("wrong access")
	}

	if !service.passwordService.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("wrong access")
	}

	return &dto.UserOutputDto{
		Id:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, err
}
