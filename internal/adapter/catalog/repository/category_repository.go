package repository

import (
	"fmt"

	"khrix/egommerce/internal/infrastructure/database/models"
	dbhelper "khrix/egommerce/internal/infrastructure/libs/db_helper"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	database          *gorm.DB
	subCategFieldName string
}

func NewCategoryRepository(database *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		database:          database,
		subCategFieldName: fmt.Sprintf("Sub%s", dbhelper.GetReflectName(&models.Category{})),
	}
}

func (repo CategoryRepository) CreateCategory(name string) (*models.Category, error) {
	newCategory := &models.Category{Name: name}
	result := repo.database.Create(newCategory)

	return newCategory, result.Error
}

func (repo CategoryRepository) CreateSubCategory(name string, categoryId uint) (*models.Category, error) {
	newSubCategory := &models.Category{Name: name, CategoryID: &categoryId}
	parent := &models.Category{Model: gorm.Model{ID: categoryId}}
	result := repo.database.Create(&newSubCategory).Preload(repo.subCategFieldName, newSubCategory.ID).Find(&parent)

	return parent, result.Error
}

func (repo CategoryRepository) SetProductCategory(productId, categoryId uint) (*models.Product, error) {
	newProductCategory := &models.Product{Category: []*models.Category{{Model: gorm.Model{ID: categoryId}}}, Model: gorm.Model{ID: productId}}

	result := repo.database.Updates(newProductCategory)

	return newProductCategory, result.Error
}

func (r CategoryRepository) ListProductsFromCategory(categoryId uint) (*models.Category, error) {
	var categ models.Category

	productRelations := []string{
		dbhelper.GetReflectName(&models.ProductImage{}),
		dbhelper.GetReflectName(&models.ProductReview{}),
		dbhelper.GetReflectName(&models.ProductFeatureItem{}),
		dbhelper.GetReflectName(&models.Category{}),
	}

	tx := r.database.Preload(dbhelper.GetReflectName(&models.Product{}))

	for _, tableName := range productRelations {
		tx.Preload(dbhelper.NestedTableName(dbhelper.GetReflectName(&models.Product{}), tableName))
	}

	result := tx.Find(&categ, categoryId)

	return &categ, result.Error
}

func (r CategoryRepository) ListAllCategories(categoryId int32) (*[]models.Category, error) {
	var parents []models.Category

	subQueryWhere := fmt.Sprintf("with recursive cte (id, name, category_id) as (SELECT id, name, category_id FROM `categories` WHERE id = %d OR category_id = %d union all SELECT c.id, c.name, c.category_id FROM `categories` c join cte on c.category_id = cte.id ) select id from cte union select category_id from cte", categoryId, categoryId)

	subQuery := "with recursive cte (id, name, category_id) as (SELECT id, name, category_id FROM `categories` union all SELECT c.id, c.name, c.category_id FROM `categories` c join cte on c.category_id = cte.id ) select id from cte union select category_id from cte"

	result := r.database.Where("id in (?)", func() *gorm.DB {
		if categoryId == 0 {
			return r.database.Raw(subQuery)
		}
		return r.database.Raw(subQueryWhere)
	}()).Find(&parents).Order("id desc")

	mapResult := make(map[uint]models.Category, 0)

	for _, item := range parents {
		mapResult[item.ID] = item
	}

	for _, item := range parents {
		isChild := item.CategoryID != nil

		if isChild {

			parent := mapResult[*item.CategoryID]

			if parent.ID == 0 {
				continue
			}

			r.recursiveAddCategorie(uint(categoryId), &item, &parents, &mapResult)

			parent.SubCategory = append(parent.SubCategory, item)

			mapResult[parent.ID] = parent
			delete(mapResult, item.ID)
		}
	}

	newList := make([]models.Category, 0, len(mapResult))

	for _, item := range mapResult {
		newList = append(newList, item)
	}

	return &newList, result.Error
}

func (r CategoryRepository) recursiveAddCategorie(categoryId uint, item *models.Category, parents *[]models.Category, mapResult *map[uint]models.Category) {
	granChild := make([]models.Category, 0)

	for _, granItem := range *parents {
		if granItem.CategoryID != nil && item.ID == *granItem.CategoryID {

			r.recursiveAddCategorie(categoryId, &granItem, parents, mapResult)
			granChild = append(granChild, granItem)

			delete(*mapResult, granItem.ID)

		}
	}

	if len(granChild) > 0 && categoryId != item.ID {
		item.SubCategory = granChild
		(*mapResult)[item.ID] = *item
	}
}
