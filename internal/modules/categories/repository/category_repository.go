package repository

import (
	dbhelper "khrix/egommerce/internal/libs/db_helper"
	"khrix/egommerce/internal/modules/categories/repository/entities"
	product_entities "khrix/egommerce/internal/modules/product/repository/entities"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	database  *gorm.DB
	tableName string
}

func NewCategoryRepository(database *gorm.DB) *CategoryRepository {
	return &CategoryRepository{database: database, tableName: dbhelper.GetEntityTableName(&entities.Category{})}
}

func (repo CategoryRepository) CreateCategory(name string) (*entities.Category, error) {
	newCategory := &entities.Category{Name: name}
	result := repo.database.Create(newCategory)

	return newCategory, result.Error
}

func (repo CategoryRepository) CreateSubCategory(name string, categoryId uint) (*entities.Category, error) {
	newSubCategory := &entities.Category{Category: []entities.Category{{Name: name}}, Model: gorm.Model{ID: categoryId}}
	result := repo.database.Create(newSubCategory)

	return newSubCategory, result.Error
}

func (repo CategoryRepository) SetProductCategory(productId, categoryId uint) (*product_entities.Product, error) {
	newProductCategory := &product_entities.Product{Category: []entities.Category{{Model: gorm.Model{ID: categoryId}}}}
	result := repo.database.Create(newProductCategory)

	return newProductCategory, result.Error
}
