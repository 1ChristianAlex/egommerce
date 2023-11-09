package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Password  string
	Name      string
	Email     string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
