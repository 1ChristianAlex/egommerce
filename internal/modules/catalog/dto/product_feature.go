package dto

type CreateProductFeatureInputDto struct {
	Name string `json:"name"  binding:"required"`
}

type CreateFeatureItemInputDto struct {
	Name             string `json:"name"  binding:"required"`
	ProductFeatureID int32  `json:"productFeatureId"  binding:"required"`
}

type CreateProductFeatureBindInputDto struct {
	FeatureIDS []int32 `json:"featureIds"  binding:"required"`
	ProductId  int32   `json:"productId"  binding:"required"`
}

type CreateFeatureItemBindInputDto struct {
	FeatureItemIDS []int32 `json:"featureItemIds"  binding:"required"`
	FeatureId      int32   `json:"featureId"  binding:"required"`
}

type ProductFeatureItemOutputDto struct {
	ID               int32  `json:"id"  binding:"required"`
	Name             string `json:"name"  binding:"required"`
	ProductFeatureID int32  `json:"productFeatureId"  binding:"required"`
}

type ProductQuery struct {
	FeatureIDS []int32 `form:"features"`
	Search     *string `form:"search"`
}

type ProductFeatureOutputDto struct {
	ID                 int32                          `json:"id"  binding:"required"`
	Name               string                         `json:"name"  binding:"required"`
	ProductFeatureItem *[]ProductFeatureItemOutputDto `json:"productFeatureItem"  binding:"required"`
}
