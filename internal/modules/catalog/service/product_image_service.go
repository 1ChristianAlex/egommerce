package service

import (
	"bufio"
	"errors"
	"mime/multipart"

	libDi "khrix/egommerce/internal/libs/file_upload/di"
	"khrix/egommerce/internal/modules/catalog/di"
	"khrix/egommerce/internal/modules/catalog/repository/entities"
)

type ProductImageService struct {
	productImageRepository di.ProductImageRepository
	productRepository      di.ProductRepository
	fileManager            libDi.FileUploadManager
}

func NewProductImageService(
	productImageRepository di.ProductImageRepository,
	productRepository di.ProductRepository,
	fileManager libDi.FileUploadManager,
) *ProductImageService {
	return &ProductImageService{
		productImageRepository: productImageRepository,
		fileManager:            fileManager,
		productRepository:      productRepository,
	}
}

func (service ProductImageService) uploadToStorage(file *multipart.FileHeader) (*string, error) {
	fileOpen, err := file.Open()
	if err != nil {
		return nil, errors.New("erro on open file upload")
	}

	reader := bufio.NewReader(fileOpen)

	var content []byte
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		// Check for end of file
		if err != nil {
			break
		}

		// Append the read bytes to the content slice
		content = append(content, buffer[:n]...)
	}

	path, uploadError := service.fileManager.UploadFile(content, file.Filename)

	if uploadError != nil {
		return nil, errors.New("error on upload to storage")
	}

	return path, nil
}

func (service ProductImageService) UploadProductImage(file *multipart.FileHeader, productId uint) error {
	path, errorUpload := service.uploadToStorage(file)

	if errorUpload != nil {
		return errors.New("error on save image")
	}

	_, saveError := service.productImageRepository.CreateNewImageProduct(&[]entities.ProductImage{{Source: *path, ProductID: &productId}})

	if saveError != nil {
		return errors.New("error on save image")
	}

	return nil
}
