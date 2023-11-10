package repository

import (
	"khrix/egommerce/internal/modules/product/repository/entities"

	"gorm.io/gorm"
)

type ProductImageRepository struct {
	database *gorm.DB
}

func NewProductImageRepository(database *gorm.DB) *ProductImageRepository {
	return &ProductImageRepository{
		database: database,
	}
}

func (repo ProductImageRepository) CreateNewImageProduct(productImageItem *[]entities.ProductImage) (*[]entities.ProductImage, error) {
	result := repo.database.Create(&productImageItem)

	return productImageItem, result.Error
}
