package di

import (
	"khrix/egommerce/internal/modules/categories/dto"
	product_dto "khrix/egommerce/internal/modules/product/dto"
)

type CategoryService interface {
	CreateCategory(name string) (*dto.CategoryOutputDto, error)
	CreateSubCategory(name string, categoryId uint) (*dto.CategoryOutputDto, error)
	SetProductCategory(productId, categoryId uint) (*product_dto.ProductOutputDto, error)
}
