package mapper

import (
	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
)

type ProductFeatureMapper struct {
	productFeatureItemMapper di.ProductFeatureItemMapper
}

func NewProductFeatureMapper(
	productFeatureItemMapper di.ProductFeatureItemMapper,
) *ProductFeatureMapper {
	return &ProductFeatureMapper{
		productFeatureItemMapper: productFeatureItemMapper,
	}
}

func (m ProductFeatureMapper) ToDto(item entities.ProductFeature) *dto.ProductFeatureOutputDto {
	featureItem := addons.Map(item.ProductFeatureItem, func(i entities.ProductFeatureItem) dto.ProductFeatureItemOutputDto {
		return *m.productFeatureItemMapper.ToDto(i)
	})

	return &dto.ProductFeatureOutputDto{
		ID:                 int32(item.ID),
		Name:               item.Name,
		ProductFeatureItem: &featureItem,
	}
}

func (m ProductFeatureMapper) ToEntity(item dto.ProductFeatureOutputDto) *entities.ProductFeature {
	return &entities.ProductFeature{
		Model: gorm.Model{ID: uint(item.ID)},
		Name:  item.Name,
		ProductFeatureItem: addons.Map(*item.ProductFeatureItem, func(i dto.ProductFeatureItemOutputDto) entities.ProductFeatureItem {
			return *m.productFeatureItemMapper.ToEntity(i)
		}),
	}
}
