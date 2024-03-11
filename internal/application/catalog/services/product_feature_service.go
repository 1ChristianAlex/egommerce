package service

import (
	"errors"

	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/addons"
	"khrix/egommerce/internal/infrastructure/database/models"
	"khrix/egommerce/internal/port/catalog"
)

type ProductFeatureService struct {
	productFeatureRepository catalog.ProductFeatureRepository
	productFeatureMapper     catalog.ProductFeatureMapper[models.ProductFeature]
	productFeatureItemMapper catalog.ProductFeatureItemMapper
	productMapper            catalog.ProductMapper[models.Product]
}

func NewProductFeatureService(
	productFeatureRepository catalog.ProductFeatureRepository,
	productFeatureMapper catalog.ProductFeatureMapper[models.ProductFeature],
	productFeatureItemMapper catalog.ProductFeatureItemMapper,
	productMapper catalog.ProductMapper[models.Product],
) *ProductFeatureService {
	return &ProductFeatureService{
		productFeatureRepository: productFeatureRepository,
		productFeatureMapper:     productFeatureMapper,
		productFeatureItemMapper: productFeatureItemMapper,
		productMapper:            productMapper,
	}
}

func (s ProductFeatureService) CreateProductFeature(featureName string) (*dto.ProductFeatureOutputDto, error) {
	feature, err := s.productFeatureRepository.CreateProductFeature(featureName)
	if err != nil {
		return nil, errors.New("error on feature creation")
	}

	result := s.productFeatureMapper.ToDto(*feature)

	return result, nil
}

func (s ProductFeatureService) CreateProductFeatureItem(featureItemName string, featureId int32) (*dto.ProductFeatureItemOutputDto, error) {
	featureItem, err := s.productFeatureRepository.CreateFeatureItem(featureItemName, uint(featureId))
	if err != nil {
		return nil, errors.New("error on feature item creation")
	}

	result := s.productFeatureItemMapper.ToDto(*featureItem)

	return result, nil
}

func (s ProductFeatureService) BindProductWithFeature(productId int32, featureIds []int32) error {
	_, err := s.productFeatureRepository.BindProductWithFeature(uint(productId), addons.Map(featureIds, func(item int32) uint { return uint(item) }))

	return err
}

func (s ProductFeatureService) BindFeatureWithItem(featureId int32, featureItemId []int32) error {
	_, err := s.productFeatureRepository.BindFeatureWithItem(uint(featureId), addons.Map(featureItemId, func(item int32) uint { return uint(item) }))

	return err
}

func (s ProductFeatureService) FindProductsByFeature(featureIds []int32) (*[]dto.ProductOutputDto, error) {
	productList, err := s.productFeatureRepository.FindProductsByFeatureItem(addons.Map(featureIds, func(item int32) uint { return uint(item) }))
	if err != nil {
		return nil, errors.New("error on finding products by feature")
	}

	result := addons.Map(productList, func(item *models.Product) dto.ProductOutputDto { return *s.productMapper.ToDto(*item) })

	return &result, nil
}
