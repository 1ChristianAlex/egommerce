package dto

type CategoryOutputDto struct {
	ID          uint                `json:"id"`
	SubCategory []CategoryOutputDto `json:"subCategory"`
	Name        string              `json:"name"`
}

type CreateCategoryInputDto struct {
	Name       string `json:"name" binding:"required"`
	CategoryId uint   `json:"categoryId"`
}

type GetProductsCategory struct {
	CategoryId int32 `uri:"categoryId" binding:"required"`
}

type SetProductCategoryInputDto struct {
	ProductId  uint `json:"productId" binding:"required"`
	CategoryId uint `json:"categoryId" binding:"required"`
}
