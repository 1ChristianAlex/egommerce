package repository

import (
	"reflect"

	"khrix/egommerce/internal/modules/product/repository/entities"

	"gorm.io/gorm"
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
	result := repo.database.Create(&productItem).Association(reflect.TypeOf(&entities.ProductImage{}).Elem().Name())

	return productItem, result.Error
}
