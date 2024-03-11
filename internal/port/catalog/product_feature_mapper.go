package catalog

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/port/libs"
)

type ProductFeatureMapper[R any] interface {
	libs.BaseMapper[dto.ProductFeatureOutputDto, dto.ProductFeatureOutputDto, R]
}
