package migrations

import (
	"fmt"

	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
)

func ProductMigration(database *gorm.DB) {
	fmt.Println("Migrating")
	database.Migrator().CurrentDatabase()

	entitiesList := []interface{}{
		entities.Category{},
		entities.Product{},
		entities.ProductFeature{},
		entities.ProductFeatureItem{},
		entities.ProductImage{},
		entities.ProductReview{},
		entities.ProductFeatureItem{},
		entities.ProductFeature{},
	}

	for _, v := range entitiesList {
		database.AutoMigrate(&v)
	}

	firstCategory := entities.Category{
		Name: "Test Category",
	}

	database.FirstOrCreate(&firstCategory)

	firstCategory.SubCategory = []entities.Category{{Name: "Sub Category", CategoryID: &firstCategory.ID}}

	database.Where(entities.Category{Model: firstCategory.Model}).Updates(&firstCategory)

	database.FirstOrCreate(&entities.Product{
		Name:          "Product Item Test",
		Description:   "Description Test",
		Price:         1245.36,
		DiscountPrice: 1245.36,
		Quantity:      154,
		ProductImage:  []entities.ProductImage{{Source: "https://teste.com.br"}},
		Category:      []*entities.Category{&firstCategory},
	})
}
