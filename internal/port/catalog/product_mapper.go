package catalog

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/infrastructure/database/models"
)

type ProductMapper interface {
	ToDto(item models.Product) *dto.ProductOutputDto
	ToEntity(item dto.ProductInputDto) models.Product
}
