package di

import "khrix/egommerce/internal/modules/user/repository/entities"

type UserRepository interface {
	CreateUser(userModel *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
