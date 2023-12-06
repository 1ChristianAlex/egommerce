package di

import "khrix/egommerce/internal/modules/catalog/repository/entities"

type CategoryRepository interface {
	CreateCategory(name string) (*entities.Category, error)
	CreateSubCategory(name string, categoryId uint) (*entities.Category, error)
	SetProductCategory(productId, categoryId uint) (*entities.Product, error)
	ListProductsFromCategory(categoryId uint) (*entities.Category, error)
	ListAllCategories() (*[]entities.Category, error)
}
