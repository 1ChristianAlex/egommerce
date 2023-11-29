package mapper

import (
	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/product/dto"
	"khrix/egommerce/internal/modules/product/repository/entities"
)

type ProductMapper struct{}

func NewProductMapper() *ProductMapper {
	return &ProductMapper{}
}

func (m ProductMapper) ToDto(item entities.Product) dto.ProductOutputDto {
	return dto.ProductOutputDto{
		ID:            item.ID,
		Name:          item.Name,
		Description:   item.Description,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		Quantity:      item.Quantity,
		Images:        addons.Map(item.ProductImage, func(image entities.ProductImage) string { return image.Source }),
	}
}

func (m ProductMapper) ToEntity(item dto.ProductInputDto) entities.Product {
	return entities.Product{
		Name:          item.Name,
		Description:   item.Description,
		Price:         item.Price,
		Quantity:      item.Quantity,
		DiscountPrice: item.DiscountPrice,
	}
}
