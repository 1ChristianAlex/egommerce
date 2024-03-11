package repository

import (
	"fmt"

	"khrix/egommerce/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SearchRepository struct {
	database *gorm.DB
}

func NewSearchRepository(database *gorm.DB) *SearchRepository {
	return &SearchRepository{
		database: database,
	}
}

func (repo SearchRepository) Search(searchValue *string, categories, features *[]int32) (*[]models.Product, error) {
	products := []models.Product{}
	tx := repo.database.Preload(clause.Associations)

	if searchValue != nil {
		tx.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *searchValue))
	}

	if features != nil {
		tx.Joins(
			"join product_feature_mm pfm on id = pfm.product_id ",
		).Where("pfm.product_feature_item_id  in ?", *features)
	}

	if categories != nil {
		tx.Joins(
			"join product_category_mm pcm on id = pcm.product_id ",
		).Where("pcm.category_id  in ?", *categories)
	}

	result := tx.Find(&products)

	return &products, result.Error
}
