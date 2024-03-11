package auth

import "khrix/egommerce/internal/application/user/dto"

type JwtService interface {
	NewClains(user dto.UserOutputDto) (string, error)
	FromClains(tokenString string) (*dto.UserOutputTokenDto, error)
	IsValid(tokenString string) error
}
