package controller

import (
	"net/http"

	base_controller "khrix/egommerce/internal/adapter/libs/base_controller"
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/channels"
	"khrix/egommerce/internal/port/auth"
	"khrix/egommerce/internal/port/catalog"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService        catalog.ProductService
	productFeatureService catalog.ProductFeatureService
	authHelper            auth.AuthHelper
}

func NewProductController(
	productService catalog.ProductService,
	productFeatureService catalog.ProductFeatureService,
	authHelper auth.AuthHelper,
) *ProductController {
	return &ProductController{
		productService:        productService,
		productFeatureService: productFeatureService,
		authHelper:            authHelper,
	}
}

func (controller ProductController) CreateNewProductItem(context *gin.Context) {
	user, err := controller.authHelper.ExtractClaimsFromContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[interface{}]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	var productBody dto.ProductInputDto

	base_controller.ControllerInputMethod(context, &productBody, context.ShouldBindJSON, func(channel chan channels.Resolve[dto.ProductOutputDto]) {
		productResult, errProduct := controller.productService.CreateNewProduct(productBody, user.UserID)

		channel <- channels.Resolve[dto.ProductOutputDto]{Result: *productResult, Err: errProduct}
	})
}

func (controller ProductController) GetListProducts(context *gin.Context) {
	var query dto.ProductQuery

	base_controller.ControllerInputMethod(context, &query, context.ShouldBindQuery, func(channel chan channels.Resolve[[]dto.ProductOutputDto]) {
		productResult, errProduct := controller.productService.ListProducts(query.Search, query.CategoryIDS, query.FeatureIDS)
		channel <- channels.Resolve[[]dto.ProductOutputDto]{Result: *productResult, Err: errProduct}
	})
}
