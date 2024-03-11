package catalog

import "khrix/egommerce/internal/infrastructure/database/models"

type ProductFeatureRepository interface {
	CreateProductFeature(name string) (*models.ProductFeature, error)
	CreateFeatureItem(name string, pFeatureId uint) (*models.ProductFeatureItem, error)
	BindProductWithFeature(productId uint, featureId []uint) (*models.Product, error)
	BindFeatureWithItem(featureId uint, featureItem []uint) (*models.ProductFeature, error)
	FindProductsByFeatureItem(featureItemIds []uint) ([]*models.Product, error)
}
