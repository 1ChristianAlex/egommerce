package dto

type CategoryOutputDto struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	SubCategory []CategoryOutputDto `json:"subCategory"`
}

type CreateCategoryInputDto struct {
	Name       string `json:"name" binding:"required"`
	CategoryId uint   `json:"categoryId"`
}

type GetProductsCategory struct {
	CategoryId int32 `uri:"categoryId" binding:"required"`
}

type CategoryById struct {
	CategoryID int32 `form:"categoryId" `
}

type SetProductCategoryInputDto struct {
	ProductId  uint `json:"productId" binding:"required"`
	CategoryId uint `json:"categoryId" binding:"required"`
}
