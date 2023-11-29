package entities

import (
	category_entities "khrix/egommerce/internal/modules/categories/repository/entities"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string  `gorm:"not null; char"`
	Description   string  `gorm:"not null; char"`
	Price         float64 `gorm:"not null;"`
	Quantity      int32   `gorm:"not null;"`
	DiscountPrice float64 `gorm:"not null;"`
	ProductImage  []ProductImage
	Category      []category_entities.Category
}
