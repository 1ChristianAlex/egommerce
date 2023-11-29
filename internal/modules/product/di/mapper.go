package di

import (
	"khrix/egommerce/internal/modules/product/dto"
	"khrix/egommerce/internal/modules/product/repository/entities"
)

type ProductMapper interface {
	ToDto(item entities.Product) dto.ProductOutputDto
	ToEntity(item dto.ProductInputDto) entities.Product
}
