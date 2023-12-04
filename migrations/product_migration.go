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

	colorFeature := entities.ProductFeature{Name: "Color"}

	database.FirstOrCreate(&colorFeature)

	colors := []string{
		"Red",
		"Green",
		"Black",
	}

	for _, itemName := range colors {
		database.Create(&entities.ProductFeatureItem{Name: itemName, ProductFeatureID: colorFeature.ID})
	}

	firstCategory := entities.Category{
		Name: "Test Category",
	}
	database.FirstOrCreate(&firstCategory)

	for i := 0; i < 10; i++ {

		firstCategory.SubCategory = []entities.Category{{Name: fmt.Sprintf("Sub Category %d", i), CategoryID: &firstCategory.ID}}

		database.Where(entities.Category{Model: firstCategory.Model}).Updates(&firstCategory)

		database.Create(&entities.Product{
			Name:           fmt.Sprintf("Product Item Test %d", i),
			Description:    fmt.Sprintf("Description Test %d", i),
			Price:          1245.36,
			DiscountPrice:  1245.36,
			Quantity:       154,
			ProductImage:   []entities.ProductImage{{Source: "https://teste.com.br"}},
			Category:       append(addons.Map(firstCategory.SubCategory, func(item entities.Category) *entities.Category { return &item }), &firstCategory),
			UserID:         1,
			ProductFeature: []*entities.ProductFeature{&colorFeature},
		})
	}
}
