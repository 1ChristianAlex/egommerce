package di

import (
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type ProductFeatureMapper interface {
	ToDto(item entities.ProductFeature) *dto.ProductFeatureOutputDto
	ToEntity(item dto.ProductFeatureOutputDto) *entities.ProductFeature
}
