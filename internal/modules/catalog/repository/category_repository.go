package repository

import (
	dbhelper "khrix/egommerce/internal/libs/db_helper"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	newSubCategory := &entities.Category{Name: name, CategoryID: &categoryId}
	result := repo.database.Create(newSubCategory)

	return newSubCategory, result.Error
}

func (repo CategoryRepository) SetProductCategory(productId, categoryId uint) (*entities.Product, error) {
	newProductCategory := &entities.Product{Category: []*entities.Category{{Model: gorm.Model{ID: categoryId}}}, Model: gorm.Model{ID: productId}}

	result := repo.database.Updates(newProductCategory)

	return newProductCategory, result.Error
}

func (r CategoryRepository) ListProductsFromCategory(categoryId uint) (*[]entities.Product, error) {
	var productList []entities.Product

	result := r.database.Preload(clause.Associations).Where(&entities.Product{Category: []*entities.Category{{Model: gorm.Model{ID: categoryId}}}}).Find(&productList)

	return &productList, result.Error
}

func (r CategoryRepository) ListAllCategories() (*[]entities.Category, error) {
	var categoryList []entities.Category

	result := r.database.Preload(r.tableName).Find(&categoryList)

	return &categoryList, result.Error
}
