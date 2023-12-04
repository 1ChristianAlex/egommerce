package controller

import (
	"net/http"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/dto"

	"github.com/gin-gonic/gin"
)

type ProductFeatureController struct {
	productFeatureService di.ProductFeatureService
}

func NewProductFeatureController(router *gin.RouterGroup, productFeatureService di.ProductFeatureService) {
	controller := ProductFeatureController{productFeatureService: productFeatureService}

	routerGroup := router.Group("/feature")

	routerGroup.POST("/", controller.CreateProductFeature)
	routerGroup.POST("/item", controller.CreateFeatureItem)
	routerGroup.POST("/item/bind", controller.CreateFeatureItemBind)
	routerGroup.POST("/product", controller.CreateProductFeatureBind)
}

func (c ProductFeatureController) CreateProductFeature(context *gin.Context) {
	var body dto.CreateProductFeatureInputDto

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.ProductFeatureOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[dto.ProductFeatureOutputDto])
	defer close(channel)

	go func() {
		featureResult, errProduct := c.productFeatureService.CreateProductFeature(body.Name)

		channel <- models.Resolve[dto.ProductFeatureOutputDto]{Result: *featureResult, Err: errProduct}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.ProductFeatureOutputDto]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusOK, &response.ResponseResult[*dto.ProductFeatureOutputDto]{
		Result: &resolve.Result,
	})
}

func (c ProductFeatureController) CreateFeatureItem(context *gin.Context) {
	var body dto.CreateFeatureItemInputDto

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.ProductFeatureItemOutputDto]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[dto.ProductFeatureItemOutputDto])
	defer close(channel)

	go func() {
		productResult, errProduct := c.productFeatureService.CreateProductFeatureItem(body.Name, int32(body.ProductFeatureID))

		channel <- models.Resolve[dto.ProductFeatureItemOutputDto]{Result: *productResult, Err: errProduct}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[*dto.ProductFeatureItemOutputDto]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusOK, &response.ResponseResult[*dto.ProductFeatureItemOutputDto]{
		Result: &resolve.Result,
	})
}

func (c ProductFeatureController) CreateProductFeatureBind(context *gin.Context) {
	var body dto.CreateProductFeatureBindInputDto

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[interface{}]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[interface{}])
	defer close(channel)

	go func() {
		err := c.productFeatureService.BindProductWithFeature(body.ProductId, body.FeatureIDS)

		channel <- models.Resolve[interface{}]{Err: err}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[interface{}]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func (c ProductFeatureController) CreateFeatureItemBind(context *gin.Context) {
	var body dto.CreateFeatureItemBindInputDto

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[interface{}]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	channel := make(chan models.Resolve[interface{}])
	defer close(channel)

	go func() {
		err := c.productFeatureService.BindFeatureWithItem(body.FeatureId, body.FeatureItemIDS)

		channel <- models.Resolve[interface{}]{Err: err}
	}()

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &response.ResponseResult[interface{}]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
