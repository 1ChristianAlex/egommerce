package catalog

type CategoryController[C any] interface {
	CreateNewCategory(context *C)
	SetProductCategory(context *C)
	ListAllCategories(context *C)
	ListProductsFromCategory(context *C)
}
