package di

import (
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type ProductFeatureItemMapper interface {
	ToDto(item entities.ProductFeatureItem) *dto.ProductFeatureItemOutputDto
	ToEntity(item dto.ProductFeatureItemOutputDto) *entities.ProductFeatureItem
}
