package mapper

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/addons"
	"khrix/egommerce/internal/infrastructure/database/models"
	"khrix/egommerce/internal/port/catalog"

	"gorm.io/gorm"
)

type ProductFeatureMapper struct {
	productFeatureItemMapper catalog.ProductFeatureItemMapper
}

func NewProductFeatureMapper(
	productFeatureItemMapper catalog.ProductFeatureItemMapper,
) *ProductFeatureMapper {
	return &ProductFeatureMapper{
		productFeatureItemMapper: productFeatureItemMapper,
	}
}

func (m ProductFeatureMapper) ToDto(item models.ProductFeature) *dto.ProductFeatureOutputDto {
	featureItem := addons.Map(item.ProductFeatureItem, func(i models.ProductFeatureItem) dto.ProductFeatureItemOutputDto {
		return *m.productFeatureItemMapper.ToDto(i)
	})

	return &dto.ProductFeatureOutputDto{
		ID:                 int32(item.ID),
		Name:               item.Name,
		ProductFeatureItem: &featureItem,
	}
}

func (m ProductFeatureMapper) ToEntity(item dto.ProductFeatureOutputDto) *models.ProductFeature {
	return &models.ProductFeature{
		Model: gorm.Model{ID: uint(item.ID)},
		Name:  item.Name,
		ProductFeatureItem: addons.Map(*item.ProductFeatureItem, func(i dto.ProductFeatureItemOutputDto) models.ProductFeatureItem {
			return *m.productFeatureItemMapper.ToEntity(i)
		}),
	}
}
