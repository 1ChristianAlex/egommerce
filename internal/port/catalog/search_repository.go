package catalog

import "khrix/egommerce/internal/infrastructure/database/models"

type SearchRepository interface {
	Search(searchValue *string, categories, features *[]int32) (*[]models.Product, error)
}
