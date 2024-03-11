package controller

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"

	base_controller "khrix/egommerce/internal/adapter/libs/base_controller"
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/channels"
	"khrix/egommerce/internal/port/catalog"

	"github.com/gin-gonic/gin"
)

type ProductImageController struct {
	productService      catalog.ProductService
	productImageService catalog.ProductImageService
}

func NewProductImageController(productService catalog.ProductService, productImageService catalog.ProductImageService) *ProductImageController {
	return &ProductImageController{productService: productService, productImageService: productImageService}
}

func (controller ProductImageController) contextErrStatus(context *gin.Context, erroMessage string) {
	context.JSON(http.StatusBadRequest, &base_controller.ResponseResult[interface{}]{Result: nil, ErrorMessage: errors.New(erroMessage).Error()})
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
	resolver := make(chan channels.Resolve[interface{}], len(files))

	for _, file := range files {
		wg.Add(1)

		go func(currentFile *multipart.FileHeader) {
			defer wg.Done()
			uploadError := controller.productImageService.UploadProductImage(currentFile, uint(productId))

			resolver <- channels.Resolve[interface{}]{Err: uploadError}
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

	context.JSON(http.StatusOK, &base_controller.ResponseResult[dto.ProductOutputDto]{
		Result: *productAfter,
	})
}
