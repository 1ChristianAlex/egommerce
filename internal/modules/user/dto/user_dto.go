package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginInputDto struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type CreateUserInputDto struct {
	Username string    `json:"username"  binding:"required"`
	Password string    `json:"password"  binding:"required"`
	Name     string    `json:"name"  binding:"required"`
	Email    string    `json:"email"  binding:"required"`
	Birthday time.Time `json:"birthday"  binding:"required"`
}

type UserOutputDto struct {
	Username  string    `json:"username" `
	Name      string    `json:"name" `
	Email     string    `json:"email" `
	Birthday  time.Time `json:"birthday" `
	CreatedAt time.Time `json:"createdAt" `
	UpdatedAt time.Time `json:"updatedAt" `
}

type UserOutputTokenDto struct {
	UserOutputDto
	jwt.RegisteredClaims
}
