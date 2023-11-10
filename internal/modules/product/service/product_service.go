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

func (service ProductService) ListAllProducts() (*[]dto.ProductOutputDto, error) {
	productList, errList := service.productRepository.ListProducts()

	if errList != nil {
		return nil, errors.New("error on list product")
	}

	productOutputList := make([]dto.ProductOutputDto, 0)

	for _, produtItem := range *productList {

		images := make([]string, 0)

		for _, img := range produtItem.ProductImage {
			images = append(images, img.Source)
		}

		productOutputList = append(productOutputList, dto.ProductOutputDto{
			ID:            produtItem.ID,
			Name:          produtItem.Name,
			Description:   produtItem.Description,
			Price:         produtItem.Price,
			DiscountPrice: produtItem.DiscountPrice,
			Quantity:      produtItem.Quantity,
			Images:        images,
		})
	}

	return &productOutputList, nil
}
