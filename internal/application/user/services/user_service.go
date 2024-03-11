package service

import (
	"errors"

	"khrix/egommerce/internal/application/user/dto"
	"khrix/egommerce/internal/port/user"
)

type UserService struct {
	repository      user.UserRepository
	passwordService user.PasswordService
	userMapper      user.UserMapper
}

func NewUserService(repository user.UserRepository, passwordService user.PasswordService, userMapper user.UserMapper) *UserService {
	return &UserService{
		repository:      repository,
		passwordService: passwordService,
		userMapper:      userMapper,
	}
}

func (service UserService) CreateNewUser(userInput *dto.UserInputDto) (newUser *dto.UserOutputDto, error error) {
	userExist, _ := service.repository.FindByEmail(userInput.Email)

	if userExist != nil {
		return nil, errors.New("email already exist")
	}

	hash, hashError := service.passwordService.HashPassword(userInput.Password)

	if hashError != nil {
		return nil, errors.New("fail on create user password")
	}

	userModel := service.userMapper.ToEntity(*userInput)
	userModel.Password = hash

	user, err := service.repository.CreateUser(&userModel)

	result := service.userMapper.ToDto(*user)

	return &result, err
}

func (service UserService) TryLogin(email, password string) (newUser *dto.UserOutputDto, error error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("wrong access")
	}

	if !service.passwordService.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("wrong access")
	}

	result := service.userMapper.ToDto(*user)

	return &result, err
}
