package catalog

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/port/libs"
)

type CategoryMapper[R any] interface {
	libs.BaseMapper[dto.CategoryOutputDto, dto.CategoryOutputDto, R]
}
