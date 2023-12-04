package entities

import "gorm.io/gorm"

type ProductFeatureItem struct {
	gorm.Model
	Name             string `gorm:"not null; char"`
	ProductFeatureID uint
}
