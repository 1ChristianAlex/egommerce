package catalog

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/port/libs"
)

type ProductMapper[R any] interface {
	libs.BaseMapper[dto.ProductInputDto, dto.ProductOutputDto, R]
}
