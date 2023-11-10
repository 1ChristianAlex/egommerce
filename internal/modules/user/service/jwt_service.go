package service

import (
	"errors"
	"time"

	"khrix/egommerce/internal/modules/user/dto"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	signingKey []byte
}

func NewJwtService() *JwtService {
	return &JwtService{
		signingKey: []byte("SECRET"),
	}
}

func (service *JwtService) getToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.UserOutputTokenDto{}, func(token *jwt.Token) (interface{}, error) {
		return service.signingKey, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, errors.New("fail on get token info")
	}

	return token, nil
}

func (service *JwtService) NewClains(user dto.UserOutputDto) (string, error) {
	claims := dto.UserOutputTokenDto{
		UserOutputDto: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	tokenSing := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenSing.SignedString(service.signingKey)

	return token, err
}

func (service *JwtService) FromClains(tokenString string) (*dto.UserOutputTokenDto, error) {
	token, err := service.getToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.UserOutputTokenDto); ok {
		return claims, nil
	}

	return nil, errors.New("fail on get from claim")
}

func (service *JwtService) IsValid(tokenString string) error {
	token, err := service.getToken(tokenString)
	if err != nil || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
