package controller

import (
	"net/http"

	base_controller "khrix/egommerce/internal/adapter/libs/base_controller"
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/channels"
	"khrix/egommerce/internal/port/catalog"

	"github.com/gin-gonic/gin"
)

type ProductFeatureController struct {
	productFeatureService catalog.ProductFeatureService
}

func NewProductFeatureController(productFeatureService catalog.ProductFeatureService) *ProductFeatureController {
	return &ProductFeatureController{productFeatureService: productFeatureService}
}

func (c ProductFeatureController) CreateProductFeature(context *gin.Context) {
	var body dto.CreateProductFeatureInputDto

	base_controller.ControllerInputMethod(context, &body, context.ShouldBindJSON, func(channel chan channels.Resolve[dto.ProductFeatureOutputDto]) {
		featureResult, errProduct := c.productFeatureService.CreateProductFeature(body.Name)

		channel <- channels.Resolve[dto.ProductFeatureOutputDto]{Result: *featureResult, Err: errProduct}
	})
}

func (c ProductFeatureController) CreateFeatureItem(context *gin.Context) {
	var body dto.CreateFeatureItemInputDto

	base_controller.ControllerInputMethod(context, &body, context.ShouldBindJSON, func(channel chan channels.Resolve[dto.ProductFeatureItemOutputDto]) {
		productResult, errProduct := c.productFeatureService.CreateProductFeatureItem(body.Name, int32(body.ProductFeatureID))

		channel <- channels.Resolve[dto.ProductFeatureItemOutputDto]{Result: *productResult, Err: errProduct}
	})
}

func (c ProductFeatureController) CreateProductFeatureBind(context *gin.Context) {
	var body dto.CreateProductFeatureBindInputDto

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[interface{}]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan channels.Resolve[interface{}])
	defer close(channel)

	go func() {
		err := c.productFeatureService.BindProductWithFeature(body.ProductId, body.FeatureIDS)

		channel <- channels.Resolve[interface{}]{Err: err}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[interface{}]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func (c ProductFeatureController) CreateFeatureItemBind(context *gin.Context) {
	var body dto.CreateFeatureItemBindInputDto

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[interface{}]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan channels.Resolve[interface{}])
	defer close(channel)

	go func() {
		err := c.productFeatureService.BindFeatureWithItem(body.FeatureId, body.FeatureItemIDS)

		channel <- channels.Resolve[interface{}]{Err: err}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[interface{}]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
