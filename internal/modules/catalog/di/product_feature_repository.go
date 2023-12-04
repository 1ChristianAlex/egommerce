package di

import "khrix/egommerce/internal/modules/catalog/repository/entities"

type ProductFeatureRepository interface {
	CreateProductFeature(name string) (*entities.ProductFeature, error)
	CreateFeatureItem(name string, pFeatureId uint) (*entities.ProductFeatureItem, error)
	BindProductWithFeature(productId uint, featureId []uint) (*entities.Product, error)
	BindFeatureWithItem(featureId uint, featureItem []uint) (*entities.ProductFeature, error)
	FindProductsByFeatureItem(featureItemIds []uint) ([]*entities.Product, error)
}
