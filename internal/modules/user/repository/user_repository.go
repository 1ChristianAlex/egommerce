package repository

import (
	"khrix/egommerce/internal/modules/user/repository/entities"

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

func (repo *UserRepository) CreateUser(userModel *entities.User) (int64, error) {
	result := repo.database.Create(&userModel)

	return result.RowsAffected, result.Error
}
