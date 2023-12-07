package dto

type ProductInputDto struct {
	Name          string  `json:"name"  binding:"required"`
	Description   string  `json:"description"  binding:"required"`
	Price         float64 `json:"price"  binding:"required"`
	DiscountPrice float64 `json:"discountPrice"  binding:"required"`
	Quantity      int32   `json:"quantity"  binding:"required"`
}

type ProductOutputDto struct {
	ID            uint                          `json:"id"  binding:"required"`
	Name          string                        `json:"name"  binding:"required"`
	Description   string                        `json:"description"  binding:"required"`
	Price         float64                       `json:"price"  binding:"required"`
	DiscountPrice float64                       `json:"discountPrice"  binding:"required"`
	Quantity      int32                         `json:"quantity"  binding:"required"`
	Images        []string                      `json:"images"  binding:"required"`
	Category      []CategoryOutputDto           `json:"category"  binding:"required"`
	Feature       []ProductFeatureItemOutputDto `json:"feature"  binding:"required"`
	UserId        int32                         `json:"userId"  binding:"required"`
}
