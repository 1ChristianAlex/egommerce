package catalog

type ProductController[C any] interface {
	CreateNewProductItem(context *C)
	GetListProducts(context *C)
}
