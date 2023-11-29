package mapper

import (
	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type ProductMapper struct {
	categoryMapper di.CategoryMapper
}

func NewProductMapper(categoryMapper di.CategoryMapper) *ProductMapper {
	return &ProductMapper{
		categoryMapper: categoryMapper,
	}
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
		Category:      addons.Map(item.Category, func(image *entities.Category) dto.CategoryOutputDto { return m.categoryMapper.ToDto(*image) }),
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
