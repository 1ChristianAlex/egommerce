package controller

import (
	"net/http"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	"khrix/egommerce/internal/modules/categories/di"
	"khrix/egommerce/internal/modules/categories/dto"
	product_dto "khrix/egommerce/internal/modules/product/dto"

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
}

func (c CategoryController) CreateNewCategory(context *gin.Context) {
	var categoryBody dto.CreateCategoryInputDto

	if err := context.ShouldBindJSON(&categoryBody); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.CategoryOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[dto.CategoryOutputDto])
	defer close(channel)

	go func() {
		if categoryBody.CategoryId == 0 {
			newCategory, err := c.categoryService.CreateCategory(categoryBody.Name)
			channel <- models.Resolve[dto.CategoryOutputDto]{Result: *newCategory, Err: err}
		} else {
			newCategory, err := c.categoryService.CreateSubCategory(categoryBody.Name, categoryBody.CategoryId)
			channel <- models.Resolve[dto.CategoryOutputDto]{Result: *newCategory, Err: err}
		}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.CategoryOutputDto]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusOK, &response.ResponseResult[*dto.CategoryOutputDto]{
		Result: &resolve.Result,
	})
}

func (c CategoryController) SetProductCategory(context *gin.Context) {
	var productCategory dto.SetProductCategoryInputDto

	if err := context.ShouldBindJSON(&productCategory); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*product_dto.ProductOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[product_dto.ProductOutputDto])
	defer close(channel)

	go func() {
		productItem, err := c.categoryService.SetProductCategory(productCategory.ProductId, productCategory.CategoryId)
		channel <- models.Resolve[product_dto.ProductOutputDto]{Result: *productItem, Err: err}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*product_dto.ProductOutputDto]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusOK, &response.ResponseResult[*product_dto.ProductOutputDto]{
		Result: &resolve.Result,
	})
}
