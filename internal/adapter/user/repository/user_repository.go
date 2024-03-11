package repository

import (
	"errors"

	"khrix/egommerce/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (repo *UserRepository) CreateUser(userModel *models.User) (*models.User, error) {
	result := repo.database.Create(&userModel)

	return userModel, result.Error
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := models.User{Email: email}

	result := repo.database.Where(&user).First(&user)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("wrong access")
	}

	return &user, result.Error
}
