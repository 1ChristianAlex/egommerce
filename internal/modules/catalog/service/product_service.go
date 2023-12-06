package service

import (
	"errors"

	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type ProductService struct {
	productRepository      di.ProductRepository
	searchRepository       di.SearchRepository
	productImageRepository di.ProductImageRepository
	productMapper          di.ProductMapper
}

func NewProductService(
	productRepository di.ProductRepository,
	searchRepository di.SearchRepository,
	productImageRepository di.ProductImageRepository,
	productMapper di.ProductMapper,
) *ProductService {
	return &ProductService{
		productRepository:      productRepository,
		productImageRepository: productImageRepository,
		productMapper:          productMapper,
		searchRepository:       searchRepository,
	}
}

func (service ProductService) CreateNewProduct(productItem dto.ProductInputDto) (*dto.ProductOutputDto, error) {
	entityItem := service.productMapper.ToEntity(productItem)
	newProduct, productErr := service.productRepository.CreateNewProduct(&entityItem)

	if productErr != nil {
		return nil, errors.New("error on create product")
	}

	mapped := service.productMapper.ToDto(*newProduct)

	return mapped, nil
}

func (service ProductService) ListProducts(searchValue *string, categories, features *[]int32) (*[]dto.ProductOutputDto, error) {
	productList, errList := service.searchRepository.Search(searchValue, categories, features)

	if errList != nil {
		return nil, errors.New("error on list product")
	}

	productOutputList := addons.Map(*productList, func(produtItem entities.Product) dto.ProductOutputDto {
		return *service.productMapper.ToDto(produtItem)
	})

	return &productOutputList, nil
}

func (service ProductService) FindById(productId uint) (*dto.ProductOutputDto, error) {
	productItem, errList := service.productRepository.FindById(productId)

	if errList != nil {
		return nil, errors.New("error on list product")
	}

	mapped := service.productMapper.ToDto(*productItem)

	return mapped, nil
}
