package catalog

import (
	"mime/multipart"

	"khrix/egommerce/internal/application/catalog/dto"
)

type (
	ProductService interface {
		CreateNewProduct(productItem dto.ProductInputDto, userId int32) (*dto.ProductOutputDto, error)
		ListProducts(searchValue *string, categories, features *[]int32) (*[]dto.ProductOutputDto, error)
		FindById(productId uint) (*dto.ProductOutputDto, error)
	}

	ProductImageService interface {
		UploadProductImage(file *multipart.FileHeader, productId uint) error
	}
)
