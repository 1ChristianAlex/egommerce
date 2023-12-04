package repository

import (
	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductFeatureRepository struct {
	database *gorm.DB
}

func NewProductFeatureRepository(database *gorm.DB) *ProductFeatureRepository {
	return &ProductFeatureRepository{database: database}
}

func (repo ProductFeatureRepository) CreateProductFeature(name string) (*entities.ProductFeature, error) {
	newCategory := &entities.ProductFeature{Name: name}
	result := repo.database.Create(newCategory)

	return newCategory, result.Error
}

func (repo ProductFeatureRepository) CreateFeatureItem(name string, pFeatureId uint) (*entities.ProductFeatureItem, error) {
	newCategory := &entities.ProductFeatureItem{Name: name, ProductFeatureID: pFeatureId}
	result := repo.database.Create(newCategory)

	return newCategory, result.Error
}

func (repo ProductFeatureRepository) BindProductWithFeature(productId uint, featureId []uint) (*entities.Product, error) {
	featureList := addons.Map(featureId, func(item uint) *entities.ProductFeature {
		return &entities.ProductFeature{Model: gorm.Model{ID: item}}
	})

	productWithFeature := &entities.Product{Model: gorm.Model{ID: productId}, ProductFeature: featureList}

	result := repo.database.Where(&entities.Product{Model: productWithFeature.Model}).Updates(productWithFeature)

	return productWithFeature, result.Error
}

func (repo ProductFeatureRepository) BindFeatureWithItem(featureId uint, featureItem []uint) (*entities.ProductFeature, error) {
	featureItemList := addons.Map(featureItem, func(item uint) entities.ProductFeatureItem {
		return entities.ProductFeatureItem{Model: gorm.Model{ID: item}}
	})

	featureUpdated := &entities.ProductFeature{Model: gorm.Model{ID: featureId}, ProductFeatureItem: featureItemList}

	result := repo.database.Where(&entities.ProductFeature{Model: featureUpdated.Model}).Updates(featureUpdated)

	return featureUpdated, result.Error
}

func (repo ProductFeatureRepository) FindProductsByFeatureItem(featureItemIds []uint) ([]*entities.Product, error) {
	finded := []*entities.Product{}

	featureItem := addons.Map(featureItemIds, func(itemId uint) *entities.ProductFeatureItem {
		return &entities.ProductFeatureItem{Model: gorm.Model{ID: itemId}}
	})

	result := repo.database.Preload(clause.Associations).Where(featureItem).Find(&finded)

	return finded, result.Error
}
