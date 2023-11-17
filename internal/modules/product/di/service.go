package di

import (
	"mime/multipart"

	"khrix/egommerce/internal/modules/product/dto"
)

type ProductService interface {
	CreateNewProduct(productItem dto.CreateProductInputDto) (*dto.ProductOutputDto, error)
	ListAllProducts() (*[]dto.ProductOutputDto, error)
	FindById(productId uint) (*dto.ProductOutputDto, error)
}

type ProductImageService interface {
	UploadProductImage(file *multipart.FileHeader, productId uint) error
}
