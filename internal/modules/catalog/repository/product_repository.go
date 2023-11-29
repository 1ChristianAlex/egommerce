package repository

import (
	"reflect"

	"khrix/egommerce/internal/modules/catalog/repository/entities"

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

func (repo ProductRepository) ListProducts() (*[]entities.Product, error) {
	productList := make([]entities.Product, 0)

	result := repo.database.Preload(reflect.TypeOf(&entities.ProductImage{}).Elem().Name()).Find(&productList)

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

	result := repo.database.Preload(reflect.TypeOf(&entities.ProductImage{}).Elem().Name()).Where(&entities.Product{Model: gorm.Model{ID: productId}}).Find(&finded)

	return &finded, result.Error
}
