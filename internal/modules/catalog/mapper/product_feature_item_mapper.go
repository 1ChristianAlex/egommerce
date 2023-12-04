package mapper

import (
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
)

type FeatureItemMapper struct{}

func NewFeatureItemMapper() *FeatureItemMapper {
	return &FeatureItemMapper{}
}

func (m FeatureItemMapper) ToDto(item entities.ProductFeatureItem) *dto.ProductFeatureItemOutputDto {
	return &dto.ProductFeatureItemOutputDto{ID: int32(item.ID), Name: item.Name, ProductFeatureID: int32(item.ProductFeatureID)}
}

func (m FeatureItemMapper) ToEntity(item dto.ProductFeatureItemOutputDto) *entities.ProductFeatureItem {
	return &entities.ProductFeatureItem{Model: gorm.Model{ID: uint(item.ID)}, Name: item.Name, ProductFeatureID: uint(item.ProductFeatureID)}
}
