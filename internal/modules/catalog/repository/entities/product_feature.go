package entities

import "gorm.io/gorm"

type ProductFeature struct {
	gorm.Model
	Name               string `gorm:"not null; char"`
	ProductFeatureItem []ProductFeatureItem
}
