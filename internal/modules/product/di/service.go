package di

import "khrix/egommerce/internal/modules/product/dto"

type ProductService interface {
	CreateNewProduct(productItem dto.CreateProductInputDto) (*dto.ProductOutputDto, error)
	ListAllProducts() (*[]dto.ProductOutputDto, error)
}
