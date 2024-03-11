package repository

import (
	"khrix/egommerce/internal/application/libs/addons"
	"khrix/egommerce/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductFeatureRepository struct {
	database *gorm.DB
}

func NewProductFeatureRepository(database *gorm.DB) *ProductFeatureRepository {
	return &ProductFeatureRepository{database: database}
}

func (repo ProductFeatureRepository) CreateProductFeature(name string) (*models.ProductFeature, error) {
	newCategory := &models.ProductFeature{Name: name}
	result := repo.database.Create(newCategory)

	return newCategory, result.Error
}

func (repo ProductFeatureRepository) CreateFeatureItem(name string, pFeatureId uint) (*models.ProductFeatureItem, error) {
	newCategory := &models.ProductFeatureItem{Name: name, ProductFeatureID: pFeatureId}
	result := repo.database.Create(newCategory)

	return newCategory, result.Error
}

func (repo ProductFeatureRepository) BindProductWithFeature(productId uint, featureItemId []uint) (*models.Product, error) {
	featureList := addons.Map(featureItemId, func(item uint) *models.ProductFeatureItem {
		return &models.ProductFeatureItem{Model: gorm.Model{ID: item}}
	})

	productWithFeature := &models.Product{Model: gorm.Model{ID: productId}, ProductFeatureItem: featureList}

	result := repo.database.Where(&models.Product{Model: productWithFeature.Model}).Updates(productWithFeature)

	return productWithFeature, result.Error
}

func (repo ProductFeatureRepository) BindFeatureWithItem(featureId uint, featureItem []uint) (*models.ProductFeature, error) {
	featureItemList := addons.Map(featureItem, func(item uint) models.ProductFeatureItem {
		return models.ProductFeatureItem{Model: gorm.Model{ID: item}}
	})

	featureUpdated := &models.ProductFeature{Model: gorm.Model{ID: featureId}, ProductFeatureItem: featureItemList}

	result := repo.database.Where(&models.ProductFeature{Model: featureUpdated.Model}).Updates(featureUpdated)

	return featureUpdated, result.Error
}

func (repo ProductFeatureRepository) FindProductsByFeatureItem(featureItemIds []uint) ([]*models.Product, error) {
	finded := []*models.Product{}

	result := repo.database.Joins(
		"join product_feature_mm pfm on id = pfm.product_id ",
	).Where("pfm.product_feature_item_id  in ?", featureItemIds).
		Preload(clause.Associations).
		Find(&finded)

	return finded, result.Error
}
