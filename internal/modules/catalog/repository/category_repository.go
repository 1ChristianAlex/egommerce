package repository

import (
	"fmt"

	dbhelper "khrix/egommerce/internal/libs/db_helper"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	database *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) *CategoryRepository {
	return &CategoryRepository{database: database}
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

func (r CategoryRepository) ListProductsFromCategory(categoryId uint) (*entities.Category, error) {
	var categ entities.Category

	productRelations := []string{
		dbhelper.GetReflectName(&entities.ProductImage{}),
		dbhelper.GetReflectName(&entities.ProductReview{}),
		dbhelper.GetReflectName(&entities.ProductFeatureItem{}),
		dbhelper.GetReflectName(&entities.Category{}),
	}

	tx := r.database.Preload(dbhelper.GetReflectName(&entities.Product{}))

	for _, tableName := range productRelations {
		tx.Preload(dbhelper.NestedTableName(dbhelper.GetReflectName(&entities.Product{}), tableName))
	}

	result := tx.Find(&categ, categoryId)

	return &categ, result.Error
}

func (r CategoryRepository) ListAllCategories() (*[]entities.Category, error) {
	var categoryList []entities.Category

	result := r.database.Preload(fmt.Sprintf("Sub%s", dbhelper.GetReflectName(&entities.Category{}))).Find(&categoryList)

	return &categoryList, result.Error
}
