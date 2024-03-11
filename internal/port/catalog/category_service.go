package catalog

import "khrix/egommerce/internal/application/catalog/dto"

type CategoryService interface {
	CreateCategory(name string) (*dto.CategoryOutputDto, error)
	CreateSubCategory(name string, categoryId uint) (*dto.CategoryOutputDto, error)
	SetProductCategory(productId, categoryId uint) (*dto.ProductOutputDto, error)
	ListAllCategories(categoryId int32) (*[]dto.CategoryOutputDto, error)
	ProductsFromCategory(cagoryId int32) (*[]dto.ProductOutputDto, error)
}
