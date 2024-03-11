package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginInputDto struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserInputDto struct {
	ID       uint      `json:"id" `
	Username string    `json:"username"  binding:"required"`
	Password string    `json:"password"  binding:"required"`
	Name     string    `json:"name"  binding:"required"`
	Email    string    `json:"email"  binding:"required"`
	Birthday time.Time `json:"birthday"  binding:"required"`
}

type UserOutputDto struct {
	UserID   int32     `json:"id" `
	Username string    `json:"username" `
	Name     string    `json:"name" `
	Email    string    `json:"email" `
	Birthday time.Time `json:"birthday" `
}

type UserOutputTokenDto struct {
	UserOutputDto
	jwt.RegisteredClaims
}
