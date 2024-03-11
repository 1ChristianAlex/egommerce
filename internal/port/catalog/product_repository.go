package catalog

import (
	"khrix/egommerce/internal/infrastructure/database/models"
)

type (
	ProductRepository interface {
		CreateNewProduct(productItem *models.Product) (*models.Product, error)
		ListProducts() (*[]models.Product, error)
		UpdateProductItem(productItem *models.Product) (*models.Product, error)
		DeleteProductItem(productId uint) error
		FindById(productId uint) (*models.Product, error)
		FindByName(productName string) (*[]models.Product, error)
	}

	ProductImageRepository interface {
		CreateNewImageProduct(productImageItem *[]models.ProductImage) (*[]models.ProductImage, error)
	}
)
