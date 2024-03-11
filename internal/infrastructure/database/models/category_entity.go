package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string     `gorm:"not null; char"`
	Product     []*Product `gorm:"many2many:product_category_mm;"`
	SubCategory []Category `gorm:"foreignkey:CategoryID; default:null"`
	CategoryID  *uint
}
