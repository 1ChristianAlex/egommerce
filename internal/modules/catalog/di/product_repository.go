package di

import (
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type (
	ProductRepository interface {
		CreateNewProduct(productItem *entities.Product) (*entities.Product, error)
		ListProducts() (*[]entities.Product, error)
		UpdateProductItem(productItem *entities.Product) (*entities.Product, error)
		DeleteProductItem(productId uint) error
		FindById(productId uint) (*entities.Product, error)
		FindByName(productName string) (*[]entities.Product, error)
	}

	ProductImageRepository interface {
		CreateNewImageProduct(productImageItem *[]entities.ProductImage) (*[]entities.ProductImage, error)
	}
)
