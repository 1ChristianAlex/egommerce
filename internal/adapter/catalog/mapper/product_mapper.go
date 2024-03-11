package mapper

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/addons"
	"khrix/egommerce/internal/infrastructure/database/models"
	"khrix/egommerce/internal/port/catalog"
)

type ProductMapper struct {
	categoryMapper catalog.CategoryMapper
	featureMapper  catalog.ProductFeatureItemMapper
}

func NewProductMapper(
	categoryMapper catalog.CategoryMapper,
	featureMapper catalog.ProductFeatureItemMapper,
) *ProductMapper {
	return &ProductMapper{
		categoryMapper: categoryMapper,
		featureMapper:  featureMapper,
	}
}

func (m ProductMapper) ToDto(item models.Product) *dto.ProductOutputDto {
	return &dto.ProductOutputDto{
		ID:            item.ID,
		Name:          item.Name,
		Description:   item.Description,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		Quantity:      item.Quantity,
		Images:        addons.Map(item.ProductImage, func(image models.ProductImage) string { return image.Source }),
		Category: addons.Map(item.Category, func(image *models.Category) dto.CategoryOutputDto {
			return *m.categoryMapper.ToDto(*image)
		}),
		Feature: addons.Map(item.ProductFeatureItem, func(item *models.ProductFeatureItem) dto.ProductFeatureItemOutputDto {
			return *m.featureMapper.ToDto(*item)
		}),
		UserId: int32(item.UserID),
	}
}

func (m ProductMapper) ToEntity(item dto.ProductInputDto) models.Product {
	return models.Product{
		Name:          item.Name,
		Description:   item.Description,
		Price:         item.Price,
		Quantity:      item.Quantity,
		DiscountPrice: item.DiscountPrice,
	}
}
