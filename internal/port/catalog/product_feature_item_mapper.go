package catalog

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/infrastructure/database/models"
)

type ProductFeatureItemMapper interface {
	ToDto(item models.ProductFeatureItem) *dto.ProductFeatureItemOutputDto
	ToEntity(item dto.ProductFeatureItemOutputDto) *models.ProductFeatureItem
}
