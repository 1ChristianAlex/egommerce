package di

import "khrix/egommerce/internal/modules/user/dto"

type JwtService interface {
	NewClains(user dto.UserOutputDto) (string, error)
	FromClains(tokenString string) (*dto.UserOutputTokenDto, error)
	IsValid(tokenString string) error
}
