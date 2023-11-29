package repository

import (
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
	result := repo.database.Create(&productItem).Association(dbhelper.GetEntityTableName(&entities.ProductImage{}))

	return productItem, result.Error
}

func (repo ProductRepository) ListProducts() (*[]entities.Product, error) {
	productList := make([]entities.Product, 0)

	result := repo.database.Preload(dbhelper.GetEntityTableName(&entities.ProductImage{})).Preload(dbhelper.GetEntityTableName(&entities.Category{})).Find(&productList)

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
