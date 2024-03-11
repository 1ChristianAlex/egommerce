package repository

import (
	"fmt"

	"khrix/egommerce/internal/infrastructure/database/models"
	dbhelper "khrix/egommerce/internal/infrastructure/libs/db_helper"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	database *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{
		database: database,
	}
}

func (repo ProductRepository) CreateNewProduct(productItem *models.Product) (*models.Product, error) {
	result := repo.database.Create(&productItem).Find(&productItem)

	return productItem, result.Error
}

func (repo ProductRepository) ListProducts() (*[]models.Product, error) {
	productList := make([]models.Product, 0)

	result := repo.database.Preload(
		dbhelper.GetReflectName(&models.ProductImage{}),
	).Preload(
		dbhelper.GetReflectName(&models.Category{}),
	).Preload(
		dbhelper.GetReflectName(&models.ProductFeature{}),
	).Find(&productList)

	return &productList, result.Error
}

func (repo ProductRepository) UpdateProductItem(productItem *models.Product) (*models.Product, error) {
	result := repo.database.Updates(&productItem)

	return productItem, result.Error
}

func (repo ProductRepository) DeleteProductItem(productId uint) error {
	result := repo.database.Delete(&models.Product{Model: gorm.Model{ID: productId}})

	return result.Error
}

func (repo ProductRepository) FindById(productId uint) (*models.Product, error) {
	finded := models.Product{}

	result := repo.database.Preload(clause.Associations).Where(&models.Product{Model: gorm.Model{ID: productId}}).Find(&finded)

	return &finded, result.Error
}

func (repo ProductRepository) FindByName(productName string) (*[]models.Product, error) {
	products := []models.Product{}

	result := repo.database.Preload(clause.Associations).Where("name LIKE ?", fmt.Sprintf("%%%s%%", productName)).Find(&products)

	return &products, result.Error
}
