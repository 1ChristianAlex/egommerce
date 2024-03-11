package catalog

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/infrastructure/database/models"
)

type ProductFeatureMapper interface {
	ToDto(item models.ProductFeature) *dto.ProductFeatureOutputDto
	ToEntity(item dto.ProductFeatureOutputDto) *models.ProductFeature
}
