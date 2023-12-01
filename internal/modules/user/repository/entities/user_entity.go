package entities

import (
	"time"

	product_entity "khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string    `gorm:"not null; char"`
	Password      string    `gorm:"not null; char"`
	Name          string    `gorm:"not null; char"`
	Email         string    `gorm:"not null; unique; char"`
	Birthday      time.Time `gorm:"not null"`
	ProductReview []product_entity.ProductReview
	Product       []product_entity.Product
}
