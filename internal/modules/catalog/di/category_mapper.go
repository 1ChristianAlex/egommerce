package di

import (
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type CategoryMapper interface {
	ToDto(item entities.Category) *dto.CategoryOutputDto
	ToEntity(item dto.CategoryOutputDto) *entities.Category
}
