package catalog

type ProductImageController[C any] interface {
	UploadImage(context *C)
}
