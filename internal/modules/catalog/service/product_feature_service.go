package service

import (
	"errors"

	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type ProductFeatureService struct {
	productFeatureRepository di.ProductFeatureRepository
	productFeatureMapper     di.ProductFeatureMapper
	productFeatureItemMapper di.ProductFeatureItemMapper
	productMapper            di.ProductMapper
}

func NewProductFeatureService(
	productFeatureRepository di.ProductFeatureRepository,
	productFeatureMapper di.ProductFeatureMapper,
	productFeatureItemMapper di.ProductFeatureItemMapper,
	productMapper di.ProductMapper,
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

	result := addons.Map(productList, func(item *entities.Product) dto.ProductOutputDto { return s.productMapper.ToDto(*item) })

	return &result, nil
}
