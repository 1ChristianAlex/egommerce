package entities

import "gorm.io/gorm"

type ProductFeatureItem struct {
	gorm.Model
	Name             string     `gorm:"not null; char"`
	Product          []*Product `gorm:"many2many:product_feature_mm;"`
	ProductFeatureID uint
}

type ProductFeatureMM struct {
	ProductID            uint `gorm:"primarykey"`
	ProductFeatureItemID uint `gorm:"primarykey"`
}
