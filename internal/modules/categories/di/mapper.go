package di

import (
	"khrix/egommerce/internal/modules/categories/dto"
	"khrix/egommerce/internal/modules/categories/repository/entities"
)

type CategoryMapper interface {
	ToDto(item entities.Category) dto.CategoryOutputDto
	ToEntity(item dto.CategoryOutputDto) entities.Category
}
