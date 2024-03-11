package user

import "khrix/egommerce/internal/infrastructure/database/models"

type UserRepository interface {
	CreateUser(userModel *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}
