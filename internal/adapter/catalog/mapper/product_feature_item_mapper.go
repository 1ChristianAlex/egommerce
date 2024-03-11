package mapper

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type FeatureItemMapper struct{}

func NewFeatureItemMapper() *FeatureItemMapper {
	return &FeatureItemMapper{}
}

func (m FeatureItemMapper) ToDto(item models.ProductFeatureItem) *dto.ProductFeatureItemOutputDto {
	return &dto.ProductFeatureItemOutputDto{ID: int32(item.ID), Name: item.Name, ProductFeatureID: int32(item.ProductFeatureID)}
}

func (m FeatureItemMapper) ToEntity(item dto.ProductFeatureItemOutputDto) *models.ProductFeatureItem {
	return &models.ProductFeatureItem{Model: gorm.Model{ID: uint(item.ID)}, Name: item.Name, ProductFeatureID: uint(item.ProductFeatureID)}
}
