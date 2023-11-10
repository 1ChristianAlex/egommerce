package service

import (
	"errors"

	"khrix/egommerce/internal/modules/product/di"
	"khrix/egommerce/internal/modules/product/dto"
	"khrix/egommerce/internal/modules/product/repository/entities"
)

type ProductService struct {
	productRepository      di.ProductRepository
	productImageRepository di.ProductImageRepository
}

func NewProductService(productRepository di.ProductRepository,
	productImageRepository di.ProductImageRepository,
) *ProductService {
	return &ProductService{
		productRepository:      productRepository,
		productImageRepository: productImageRepository,
	}
}

func (service ProductService) CreateNewProduct(productItem dto.CreateProductInputDto) (*dto.ProductOutputDto, error) {
	newProduct, productErr := service.productRepository.CreateNewProduct(&entities.Product{
		Name:          productItem.Name,
		Description:   productItem.Description,
		Price:         productItem.Price,
		DiscountPrice: productItem.DiscountPrice,
		Quantity:      productItem.Quantity,
	})

	if productErr != nil {
		return nil, errors.New("error on create product")
	}

	images := make([]string, len(newProduct.ProductImage))

	for _, img := range newProduct.ProductImage {
		images = append(images, img.Source)
	}

	return &dto.ProductOutputDto{
		Name:          newProduct.Name,
		Description:   newProduct.Description,
		Price:         newProduct.Price,
		DiscountPrice: newProduct.DiscountPrice,
		Quantity:      newProduct.Quantity,
		Images:        images,
		ID:            newProduct.ID,
	}, nil
}
