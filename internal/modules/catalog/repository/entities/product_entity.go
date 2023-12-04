package entities

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string  `gorm:"not null; char"`
	Description    string  `gorm:"not null; char"`
	Price          float64 `gorm:"not null;"`
	Quantity       int32   `gorm:"not null;"`
	DiscountPrice  float64 `gorm:"not null;"`
	ProductImage   []ProductImage
	ProductReview  []ProductReview
	ProductFeature []*ProductFeature `gorm:"many2many:product_feature;"`
	Category       []*Category       `gorm:"many2many:product_category;"`
	UserID         uint              `gorm:"not null;"`
}
