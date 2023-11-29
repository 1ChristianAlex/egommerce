package di

import (
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type ProductMapper interface {
	ToDto(item entities.Product) dto.ProductOutputDto
	ToEntity(item dto.ProductInputDto) entities.Product
}
