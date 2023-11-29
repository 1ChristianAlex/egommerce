package di

import (
	"khrix/egommerce/internal/modules/categories/repository/entities"
	product_entities "khrix/egommerce/internal/modules/product/repository/entities"
)

type CategoryRepository interface {
	CreateCategory(name string) (*entities.Category, error)
	CreateSubCategory(name string, categoryId uint) (*entities.Category, error)
	SetProductCategory(productId, categoryId uint) (*product_entities.Product, error)
}
