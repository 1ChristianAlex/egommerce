package controller

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/models"
	"khrix/egommerce/internal/modules/product/di"
	"khrix/egommerce/internal/modules/product/dto"

	"github.com/gin-gonic/gin"
)

type ProductImageController struct {
	productService      di.ProductService
	productImageService di.ProductImageService
}

func NewProductImageController(router *gin.RouterGroup, productService di.ProductService, productImageService di.ProductImageService) {
	controller := ProductImageController{productService: productService, productImageService: productImageService}

	router.PUT("/product-image/:productId", controller.UploadImage)
}

func (controller ProductImageController) contextErrStatus(context *gin.Context, erroMessage string) {
	context.JSON(http.StatusBadRequest, &response.ResponseResult[interface{}]{Result: nil, ErrorMessage: errors.New(erroMessage).Error()})
}

func (controller ProductImageController) UploadImage(context *gin.Context) {
	productPathId := context.Param("productId")

	productId, err := strconv.Atoi(productPathId)

	if productPathId == "" || err != nil {
		controller.contextErrStatus(context, "product id is required")
		return
	}

	form, formErr := context.MultipartForm()

	if formErr != nil {
		controller.contextErrStatus(context, "product image to update is required")
		return
	}

	files := form.File["upload"]

	var wg sync.WaitGroup
	resolver := make(chan models.Resolve[interface{}], len(files))

	for _, file := range files {
		wg.Add(1)

		go func(currentFile *multipart.FileHeader) {
			defer wg.Done()
			uploadError := controller.productImageService.UploadProductImage(currentFile, uint(productId))

			resolver <- models.Resolve[interface{}]{Err: uploadError}
		}(file)
	}

	go func() {
		wg.Wait()
		close(resolver)
	}()

	for res := range resolver {
		if res.Err != nil {
			controller.contextErrStatus(context, res.Err.Error())
			return
		}
	}

	productAfter, findError := controller.productService.FindById(uint(productId))

	if findError != nil {
		controller.contextErrStatus(context, findError.Error())
		return
	}

	context.JSON(http.StatusOK, &response.ResponseResult[dto.ProductOutputDto]{
		Result: *productAfter,
	})
}
