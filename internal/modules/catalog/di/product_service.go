package di

import (
	"mime/multipart"

	"khrix/egommerce/internal/modules/catalog/dto"
)

type ProductService interface {
	CreateNewProduct(productItem dto.ProductInputDto) (*dto.ProductOutputDto, error)
	ListProducts(name *string) (*[]dto.ProductOutputDto, error)
	FindById(productId uint) (*dto.ProductOutputDto, error)
	FindByName(name string) (*[]dto.ProductOutputDto, error)
}

type ProductImageService interface {
	UploadProductImage(file *multipart.FileHeader, productId uint) error
}
