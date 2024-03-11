package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	Source          string `gorm:"not null; char"`
	ProductID       *uint
	ProductReviewID *uint
}
