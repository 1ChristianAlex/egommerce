package di

import (
	"mime/multipart"

	"khrix/egommerce/internal/modules/catalog/dto"
)

type ProductService interface {
	CreateNewProduct(productItem dto.ProductInputDto) (*dto.ProductOutputDto, error)
	ListProducts(searchValue *string, categories, features *[]int32) (*[]dto.ProductOutputDto, error)
	FindById(productId uint) (*dto.ProductOutputDto, error)
}

type ProductImageService interface {
	UploadProductImage(file *multipart.FileHeader, productId uint) error
}
