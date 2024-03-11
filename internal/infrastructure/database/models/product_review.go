package models

import (
	"gorm.io/gorm"
)

type ProductReview struct {
	gorm.Model
	Title        string `gorm:"not null; char"`
	Content      string `gorm:"not null; char"`
	Stars        int32  `gorm:"not null;"`
	ProductImage []ProductImage
	UserID       uint
	ProductID    uint
}
