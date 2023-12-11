package service

import (
	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepository di.CategoryRepository
	categoryMapper     di.CategoryMapper
	productRepository  di.ProductRepository
	productMapper      di.ProductMapper
}

func NewCategoryService(
	categoryRepository di.CategoryRepository,
	categoryMapper di.CategoryMapper,
	productRepository di.ProductRepository,
	productMapper di.ProductMapper,
) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
		categoryMapper:     categoryMapper,
		productRepository:  productRepository,
		productMapper:      productMapper,
	}
}

func (c CategoryService) CreateCategory(name string) (*dto.CategoryOutputDto, error) {
	newCategory, err := c.categoryRepository.CreateCategory(name)
	if err != nil {
		return nil, err
	}

	result := c.categoryMapper.ToDto(*newCategory)

	return result, nil
}

func (c CategoryService) CreateSubCategory(name string, categoryId uint) (*dto.CategoryOutputDto, error) {
	newSubCategory, err := c.categoryRepository.CreateSubCategory(name, categoryId)
	if err != nil {
		return nil, err
	}

	result := c.categoryMapper.ToDto(*newSubCategory)

	return result, nil
}

func (c CategoryService) SetProductCategory(productId, categoryId uint) (*dto.ProductOutputDto, error) {
	_, err := c.productRepository.UpdateProductItem(&entities.Product{
		Model:    gorm.Model{ID: productId},
		Category: []*entities.Category{{Model: gorm.Model{ID: categoryId}}},
	},
	)
	if err != nil {
		return nil, err
	}

	productUpdate, err := c.productRepository.FindById(productId)
	if err != nil {
		return nil, err
	}

	return c.productMapper.ToDto(*productUpdate), nil
}

func (c CategoryService) ListAllCategories(categoryId int32) (*[]dto.CategoryOutputDto, error) {
	categories, err := c.categoryRepository.ListAllCategories(categoryId)
	if err != nil {
		return nil, err
	}

	result := addons.Map(*categories, func(item entities.Category) dto.CategoryOutputDto { return *c.categoryMapper.ToDto(item) })

	return &result, err
}

func (c CategoryService) ProductsFromCategory(cagoryId int32) (*[]dto.ProductOutputDto, error) {
	category, err := c.categoryRepository.ListProductsFromCategory(uint(cagoryId))
	if err != nil {
		return nil, err
	}

	result := addons.Map(category.Product, func(item *entities.Product) dto.ProductOutputDto { return *c.productMapper.ToDto(*item) })

	return &result, err
}
