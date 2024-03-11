package repository

import (
	"khrix/egommerce/internal/infrastructure/database/models"

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

func (repo ProductImageRepository) CreateNewImageProduct(productImageItem *[]models.ProductImage) (*[]models.ProductImage, error) {
	result := repo.database.Create(&productImageItem)

	return productImageItem, result.Error
}
