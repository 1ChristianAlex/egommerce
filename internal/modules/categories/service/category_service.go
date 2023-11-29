package service

import (
	"khrix/egommerce/internal/modules/categories/di"
	"khrix/egommerce/internal/modules/categories/dto"
	"khrix/egommerce/internal/modules/categories/repository/entities"
	product_di "khrix/egommerce/internal/modules/product/di"
	product_dto "khrix/egommerce/internal/modules/product/dto"

	product_entity "khrix/egommerce/internal/modules/product/repository/entities"

	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepository di.CategoryRepository
	categoryMapper     di.CategoryMapper
	productRepository  product_di.ProductRepository
	productMapper      product_di.ProductMapper
}

func NewCategoryService(categoryRepository di.CategoryRepository,
	categoryMapper di.CategoryMapper,
	productRepository product_di.ProductRepository,
	productMapper product_di.ProductMapper,
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

	return &result, nil
}

func (c CategoryService) CreateSubCategory(name string, categoryId uint) (*dto.CategoryOutputDto, error) {
	newSubCategory, err := c.categoryRepository.CreateSubCategory(name, categoryId)
	if err != nil {
		return nil, err
	}

	result := c.categoryMapper.ToDto(*newSubCategory)

	return &result, nil
}

func (c CategoryService) SetProductCategory(productId, categoryId uint) (*product_dto.ProductOutputDto, error) {
	productUpdate, err := c.productRepository.UpdateProductItem(&product_entity.Product{
		Model:    gorm.Model{ID: productId},
		Category: []entities.Category{{Model: gorm.Model{ID: categoryId}}},
	},
	)
	if err != nil {
		return nil, err
	}

	result := c.productMapper.ToDto(*productUpdate)

	return &result, nil
}
