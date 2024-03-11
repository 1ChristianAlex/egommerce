package migrations

import (
	"fmt"

	"khrix/egommerce/internal/application/libs/addons"
	"khrix/egommerce/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

func ProductMigration(database *gorm.DB) {
	fmt.Println("Migrating")
	database.Migrator().CurrentDatabase()

	modelsList := []interface{}{
		models.Product{},
		models.Category{},
		models.ProductImage{},
		models.ProductReview{},
		models.ProductFeature{},
		models.ProductFeatureItem{},
	}

	for _, v := range modelsList {
		err := database.AutoMigrate(&v)
		if err != nil {
			fmt.Println(err)
		}
	}

	colorFeature := models.ProductFeature{Name: "Color", ProductFeatureItem: []models.ProductFeatureItem{
		{Name: "Red"},
		{Name: "Blue"},
		{Name: "Yellow"},
		{Name: "Black"},
	}}

	if exist := database.Find(&models.ProductFeature{Name: "Color"}); exist.RowsAffected == 0 {
		database.FirstOrCreate(&colorFeature)
	}

	firstCategory := models.Category{
		Name: "Test Category",
	}

	database.FirstOrCreate(&firstCategory)

	for i := 0; i < 10; i++ {

		firstCategory.SubCategory = []models.Category{{Name: fmt.Sprintf("Sub Category %d", i), CategoryID: &firstCategory.ID}}

		database.Where(models.Category{Model: firstCategory.Model}).Updates(&firstCategory)

		database.Create(&models.Product{
			Name:               fmt.Sprintf("Product Item Test %d", i),
			Description:        fmt.Sprintf("Description Test %d", i),
			Price:              1245.36,
			DiscountPrice:      1245.36,
			Quantity:           154,
			ProductImage:       []models.ProductImage{{Source: "E:\\Projects\\egommerce\\asset\\3eb51efc-973a-43e3-b1c2-103361ebb9da.jpg"}},
			Category:           append(addons.Map(firstCategory.SubCategory, func(item models.Category) *models.Category { return &item }), &firstCategory),
			UserID:             1,
			ProductFeatureItem: addons.Map(colorFeature.ProductFeatureItem, func(item models.ProductFeatureItem) *models.ProductFeatureItem { return &item }),
		})
	}

	for i := 0; i < 10; i++ {
		randomProduct := models.Product{}
		database.First(&randomProduct)

		randomProduct.ProductReview = append(randomProduct.ProductReview, models.ProductReview{
			Title:        "Test Review",
			Content:      "Optimizer hints allow to control the query optimizer to choose a certain query execution plan, GORM supports it with gorm.io/hints, e.g:",
			Stars:        4,
			UserID:       1,
			ProductImage: []models.ProductImage{{Source: "E:\\Projects\\egommerce\\asset\\3eb51efc-973a-43e3-b1c2-103361ebb9da.jpg"}},
		})

		database.Where(&models.Product{Model: randomProduct.Model}).Updates(&randomProduct)
	}
}
