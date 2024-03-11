package catalog

type ProductFeatureController[C any] interface {
	CreateProductFeature(context *C)
	CreateFeatureItem(context *C)
	CreateProductFeatureBind(context *C)
	CreateFeatureItemBind(context *C)
}
