package entities

import "gorm.io/gorm"

type ProductFeature struct {
	gorm.Model
	Name               string     `gorm:"not null; char"`
	Product            []*Product `gorm:"many2many:product_feature;"`
	ProductFeatureItem []ProductFeatureItem
}
