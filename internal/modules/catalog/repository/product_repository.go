package repository

import (
	"fmt"

	dbhelper "khrix/egommerce/internal/libs/db_helper"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

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

func (repo ProductRepository) CreateNewProduct(productItem *entities.Product) (*entities.Product, error) {
	result := repo.database.Create(&productItem).Find(&productItem)

	return productItem, result.Error
}

func (repo ProductRepository) ListProducts() (*[]entities.Product, error) {
	productList := make([]entities.Product, 0)

	result := repo.database.Preload(
		dbhelper.GetReflectName(&entities.ProductImage{}),
	).Preload(
		dbhelper.GetReflectName(&entities.Category{}),
	).Preload(
		dbhelper.GetReflectName(&entities.ProductFeature{}),
	).Find(&productList)

	return &productList, result.Error
}

func (repo ProductRepository) UpdateProductItem(productItem *entities.Product) (*entities.Product, error) {
	result := repo.database.Updates(&productItem)

	return productItem, result.Error
}

func (repo ProductRepository) DeleteProductItem(productId uint) error {
	result := repo.database.Delete(&entities.Product{Model: gorm.Model{ID: productId}})

	return result.Error
}

func (repo ProductRepository) FindById(productId uint) (*entities.Product, error) {
	finded := entities.Product{}

	result := repo.database.Preload(clause.Associations).Where(&entities.Product{Model: gorm.Model{ID: productId}}).Find(&finded)

	return &finded, result.Error
}

func (repo ProductRepository) FindByName(productName string) (*[]entities.Product, error) {
	products := []entities.Product{}

	result := repo.database.Preload(clause.Associations).Where("name LIKE ?", fmt.Sprintf("%%%s%%", productName)).Find(&products)

	return &products, result.Error
}
