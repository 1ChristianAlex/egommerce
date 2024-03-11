package service

import (
	"errors"

	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/addons"
	"khrix/egommerce/internal/infrastructure/database/models"
	"khrix/egommerce/internal/port/catalog"
)

type ProductService struct {
	productRepository      catalog.ProductRepository
	searchRepository       catalog.SearchRepository
	productImageRepository catalog.ProductImageRepository
	productMapper          catalog.ProductMapper[models.Product]
}

func NewProductService(
	productRepository catalog.ProductRepository,
	searchRepository catalog.SearchRepository,
	productImageRepository catalog.ProductImageRepository,
	productMapper catalog.ProductMapper[models.Product],
) *ProductService {
	return &ProductService{
		productRepository:      productRepository,
		productImageRepository: productImageRepository,
		productMapper:          productMapper,
		searchRepository:       searchRepository,
	}
}

func (service ProductService) CreateNewProduct(productItem dto.ProductInputDto, userId int32) (*dto.ProductOutputDto, error) {
	entityItem := service.productMapper.ToEntity(productItem)
	entityItem.UserID = uint(userId)
	newProduct, productErr := service.productRepository.CreateNewProduct(entityItem)

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

	productOutputList := addons.Map(*productList, func(produtItem models.Product) dto.ProductOutputDto {
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
