package controller

import (
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

	response.ControllerInputMethod(context, productBody, context.ShouldBindJSON, func(channel chan models.Resolve[dto.ProductOutputDto]) {
		productResult, errProduct := controller.productService.CreateNewProduct(productBody)

		channel <- models.Resolve[dto.ProductOutputDto]{Result: *productResult, Err: errProduct}
	})
}

func (controller ProductController) GetListProducts(context *gin.Context) {
	var query dto.ProductQuery

	response.ControllerInputMethod(context, query, context.ShouldBindQuery, func(channel chan models.Resolve[[]dto.ProductOutputDto]) {
		productResult, errProduct := controller.productService.ListProducts(query.Search, query.CategoryIDS, query.FeatureIDS)
		channel <- models.Resolve[[]dto.ProductOutputDto]{Result: *productResult, Err: errProduct}
	})
}
