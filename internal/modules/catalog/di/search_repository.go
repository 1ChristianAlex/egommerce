package di

import "khrix/egommerce/internal/modules/catalog/repository/entities"

type SearchRepository interface {
	Search(searchValue *string, categories, features *[]int32) (*[]entities.Product, error)
}
