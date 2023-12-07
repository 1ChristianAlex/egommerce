package controller

import (
	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService di.CategoryService
}

func NewCategoryController(router *gin.RouterGroup, categoryService di.CategoryService) {
	controller := &CategoryController{
		categoryService: categoryService,
	}

	routerGroup := router.Group("/category")

	routerGroup.POST("/create", controller.CreateNewCategory)
	routerGroup.POST("/set-product", controller.SetProductCategory)
	routerGroup.GET("/", controller.ListAllCategories)
	routerGroup.GET("/:categoryId", controller.ListProductsFromCategory)
}

func (c CategoryController) CreateNewCategory(context *gin.Context) {
	var categoryBody dto.CreateCategoryInputDto

	response.ControllerInputMethod(context, categoryBody, context.ShouldBindJSON, func(channel chan models.Resolve[dto.CategoryOutputDto]) {
		if categoryBody.CategoryId == 0 {
			newCategory, err := c.categoryService.CreateCategory(categoryBody.Name)
			channel <- models.Resolve[dto.CategoryOutputDto]{Result: *newCategory, Err: err}
		} else {
			newCategory, err := c.categoryService.CreateSubCategory(categoryBody.Name, categoryBody.CategoryId)
			channel <- models.Resolve[dto.CategoryOutputDto]{Result: *newCategory, Err: err}
		}
	})
}

func (c CategoryController) SetProductCategory(context *gin.Context) {
	var productCategory dto.SetProductCategoryInputDto

	response.ControllerInputMethod(context, productCategory, context.ShouldBindJSON, func(channel chan models.Resolve[dto.ProductOutputDto]) {
		productItem, err := c.categoryService.SetProductCategory(productCategory.ProductId, productCategory.CategoryId)
		channel <- models.Resolve[dto.ProductOutputDto]{Result: *productItem, Err: err}
	})
}

func (c CategoryController) ListAllCategories(context *gin.Context) {
	response.ControllerBaseMethod(context, func(channel chan models.Resolve[[]dto.CategoryOutputDto]) {
		categories, err := c.categoryService.ListAllCategories()
		channel <- models.Resolve[[]dto.CategoryOutputDto]{Result: *categories, Err: err}
	})
}

func (c CategoryController) ListProductsFromCategory(context *gin.Context) {
	var query dto.GetProductsCategory

	response.ControllerInputMethod(context, query, context.ShouldBindUri, func(channel chan models.Resolve[*[]dto.ProductOutputDto]) {
		categories, err := c.categoryService.ProductsFromCategory(query.CategoryId)
		channel <- models.Resolve[*[]dto.ProductOutputDto]{Result: categories, Err: err}
	})
}
