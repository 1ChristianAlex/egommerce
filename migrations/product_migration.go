package migrations

import (
	"fmt"

	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
)

func ProductMigration(database *gorm.DB) {
	fmt.Println("Migrating")
	database.Migrator().CurrentDatabase()

	entitiesList := []interface{}{
		entities.Product{},
		entities.Category{},
		entities.ProductImage{},
		entities.ProductReview{},
		entities.ProductFeature{},
		entities.ProductFeatureItem{},
	}

	for _, v := range entitiesList {
		err := database.AutoMigrate(&v)
		if err != nil {
			fmt.Println(err)
		}
	}

	colorFeature := entities.ProductFeature{Name: "Color", ProductFeatureItem: []entities.ProductFeatureItem{
		{Name: "Red"},
		{Name: "Blue"},
		{Name: "Yellow"},
		{Name: "Black"},
	}}

	if exist := database.Find(&entities.ProductFeature{Name: "Color"}); exist.RowsAffected == 0 {
		database.FirstOrCreate(&colorFeature)
	}

	firstCategory := entities.Category{
		Name: "Test Category",
	}

	database.FirstOrCreate(&firstCategory)

	for i := 0; i < 10; i++ {

		firstCategory.SubCategory = []entities.Category{{Name: fmt.Sprintf("Sub Category %d", i), CategoryID: &firstCategory.ID}}

		database.Where(entities.Category{Model: firstCategory.Model}).Updates(&firstCategory)

		database.Create(&entities.Product{
			Name:               fmt.Sprintf("Product Item Test %d", i),
			Description:        fmt.Sprintf("Description Test %d", i),
			Price:              1245.36,
			DiscountPrice:      1245.36,
			Quantity:           154,
			ProductImage:       []entities.ProductImage{{Source: "E:\\Projects\\egommerce\\asset\\3eb51efc-973a-43e3-b1c2-103361ebb9da.jpg"}},
			Category:           append(addons.Map(firstCategory.SubCategory, func(item entities.Category) *entities.Category { return &item }), &firstCategory),
			UserID:             1,
			ProductFeatureItem: addons.Map(colorFeature.ProductFeatureItem, func(item entities.ProductFeatureItem) *entities.ProductFeatureItem { return &item }),
		})
	}

	for i := 0; i < 10; i++ {
		randomProduct := entities.Product{}
		database.First(&randomProduct)

		randomProduct.ProductReview = append(randomProduct.ProductReview, entities.ProductReview{
			Title:        "Test Review",
			Content:      "Optimizer hints allow to control the query optimizer to choose a certain query execution plan, GORM supports it with gorm.io/hints, e.g:",
			Stars:        4,
			UserID:       1,
			ProductImage: []entities.ProductImage{{Source: "E:\\Projects\\egommerce\\asset\\3eb51efc-973a-43e3-b1c2-103361ebb9da.jpg"}},
		})

		database.Where(&entities.Product{Model: randomProduct.Model}).Updates(&randomProduct)
	}
}
