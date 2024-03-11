package catalog

import "khrix/egommerce/internal/application/catalog/dto"

type ProductFeatureService interface {
	CreateProductFeature(featureName string) (*dto.ProductFeatureOutputDto, error)
	CreateProductFeatureItem(featureItemName string, featureId int32) (*dto.ProductFeatureItemOutputDto, error)
	BindProductWithFeature(productId int32, featureIds []int32) error
	BindFeatureWithItem(featureId int32, featureItemId []int32) error
	FindProductsByFeature(featureIds []int32) (*[]dto.ProductOutputDto, error)
}
