package catalog

import "khrix/egommerce/internal/infrastructure/database/models"

type CategoryRepository interface {
	CreateCategory(name string) (*models.Category, error)
	CreateSubCategory(name string, categoryId uint) (*models.Category, error)
	SetProductCategory(productId, categoryId uint) (*models.Product, error)
	ListProductsFromCategory(categoryId uint) (*models.Category, error)
	ListAllCategories(categoryId int32) (*[]models.Category, error)
}
