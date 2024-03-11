package catalog

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/infrastructure/database/models"
)

type CategoryMapper interface {
	ToDto(item models.Category) *dto.CategoryOutputDto
	ToEntity(item dto.CategoryOutputDto) *models.Category
}
