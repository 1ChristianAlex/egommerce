package controller

import (
	base_controller "khrix/egommerce/internal/adapter/libs/base_controller"
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/channels"
	"khrix/egommerce/internal/port/catalog"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService catalog.CategoryService
}

func NewCategoryController(categoryService catalog.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (c CategoryController) CreateNewCategory(context *gin.Context) {
	var categoryBody dto.CreateCategoryInputDto

	base_controller.ControllerInputMethod(context, &categoryBody, context.ShouldBindJSON, func(channel chan channels.Resolve[dto.CategoryOutputDto]) {
		if categoryBody.CategoryId == 0 {
			newCategory, err := c.categoryService.CreateCategory(categoryBody.Name)
			channel <- channels.Resolve[dto.CategoryOutputDto]{Result: *newCategory, Err: err}
		} else {
			newCategory, err := c.categoryService.CreateSubCategory(categoryBody.Name, categoryBody.CategoryId)
			channel <- channels.Resolve[dto.CategoryOutputDto]{Result: *newCategory, Err: err}
		}
	})
}

func (c CategoryController) SetProductCategory(context *gin.Context) {
	var productCategory dto.SetProductCategoryInputDto

	base_controller.ControllerInputMethod(context, &productCategory, context.ShouldBindJSON, func(channel chan channels.Resolve[dto.ProductOutputDto]) {
		productItem, err := c.categoryService.SetProductCategory(productCategory.ProductId, productCategory.CategoryId)
		channel <- channels.Resolve[dto.ProductOutputDto]{Result: *productItem, Err: err}
	})
}

func (c CategoryController) ListAllCategories(context *gin.Context) {
	var category dto.CategoryById

	base_controller.ControllerInputMethod(context, &category, context.ShouldBindQuery, func(channel chan channels.Resolve[[]dto.CategoryOutputDto]) {
		categories, err := c.categoryService.ListAllCategories(category.CategoryID)
		channel <- channels.Resolve[[]dto.CategoryOutputDto]{Result: *categories, Err: err}
	})
}

func (c CategoryController) ListProductsFromCategory(context *gin.Context) {
	var query dto.GetProductsCategory

	base_controller.ControllerInputMethod(context, &query, context.ShouldBindUri, func(channel chan channels.Resolve[*[]dto.ProductOutputDto]) {
		categories, err := c.categoryService.ProductsFromCategory(query.CategoryId)
		channel <- channels.Resolve[*[]dto.ProductOutputDto]{Result: categories, Err: err}
	})
}
