package controller

import (
	"net/http"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	auth_di "khrix/egommerce/internal/modules/auth/di"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService        di.ProductService
	productFeatureService di.ProductFeatureService
	authHelper            auth_di.AuthHelper
}

func NewProductController(
	router *gin.RouterGroup,
	productService di.ProductService,
	productFeatureService di.ProductFeatureService,
	authHelper auth_di.AuthHelper,
) {
	controller := ProductController{
		productService:        productService,
		productFeatureService: productFeatureService,
		authHelper:            authHelper,
	}

	router.POST("/product", controller.CreateNewProductItem)
	router.GET("/product", controller.GetListProducts)
}

func (controller ProductController) CreateNewProductItem(context *gin.Context) {
	user, err := controller.authHelper.ExtractClaimsFromContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[interface{}]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	var productBody dto.ProductInputDto

	response.ControllerInputMethod(context, &productBody, context.ShouldBindJSON, func(channel chan models.Resolve[dto.ProductOutputDto]) {
		productResult, errProduct := controller.productService.CreateNewProduct(productBody, user.UserID)

		channel <- models.Resolve[dto.ProductOutputDto]{Result: *productResult, Err: errProduct}
	})
}

func (controller ProductController) GetListProducts(context *gin.Context) {
	var query dto.ProductQuery

	response.ControllerInputMethod(context, &query, context.ShouldBindQuery, func(channel chan models.Resolve[[]dto.ProductOutputDto]) {
		productResult, errProduct := controller.productService.ListProducts(query.Search, query.CategoryIDS, query.FeatureIDS)
		channel <- models.Resolve[[]dto.ProductOutputDto]{Result: *productResult, Err: errProduct}
	})
}
