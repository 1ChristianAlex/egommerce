package entities

import "time"

type User struct {
	Id        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"not null; char"`
	Password  string    `gorm:"not null; char"`
	Name      string    `gorm:"not null; char"`
	Email     string    `gorm:"not null; unique; char"`
	Birthday  time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
