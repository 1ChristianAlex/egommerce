package controller

import (
	"net/http"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService        di.ProductService
	productFeatureService di.ProductFeatureService
}

func NewProductController(
	router *gin.RouterGroup,
	productService di.ProductService,
	productFeatureService di.ProductFeatureService,
) {
	controller := ProductController{
		productService:        productService,
		productFeatureService: productFeatureService,
	}

	router.POST("/product", controller.CreateNewProductItem)
	router.GET("/product", controller.GetListProducts)
}

func (controller ProductController) CreateNewProductItem(context *gin.Context) {
	var productBody dto.ProductInputDto

	if err := context.ShouldBindJSON(&productBody); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.ProductInputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[dto.ProductOutputDto])
	defer close(channel)

	go func() {
		productResult, errProduct := controller.productService.CreateNewProduct(productBody)

		channel <- models.Resolve[dto.ProductOutputDto]{Result: *productResult, Err: errProduct}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.ProductOutputDto]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusOK, &response.ResponseResult[*dto.ProductOutputDto]{
		Result: &resolve.Result,
	})
}

func (controller ProductController) GetListProducts(context *gin.Context) {
	var query dto.ProductQuery

	if err := context.ShouldBindQuery(&query); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.ProductOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[[]dto.ProductOutputDto])
	defer close(channel)

	go func() {
		productResult, errProduct := controller.productService.ListProducts(query.Search, query.CategoryIDS, query.FeatureIDS)
		channel <- models.Resolve[[]dto.ProductOutputDto]{Result: *productResult, Err: errProduct}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*[]dto.ProductOutputDto]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusOK, &response.ResponseResult[*[]dto.ProductOutputDto]{
		Result: &resolve.Result,
	})
}
