package di

import "khrix/egommerce/internal/modules/product/repository/entities"

type ProductRepository interface {
	CreateNewProduct(productItem *entities.Product) (*entities.Product, error)
	ListProducts() (*[]entities.Product, error)
}

type ProductImageRepository interface {
	CreateNewImageProduct(productImageItem *[]entities.ProductImage) (*[]entities.ProductImage, error)
}
