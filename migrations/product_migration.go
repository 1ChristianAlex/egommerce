package migrations

import (
	"fmt"

	"khrix/egommerce/internal/modules/product/repository/entities"

	"gorm.io/gorm"
)

func ProductMigration(database *gorm.DB) {
	fmt.Println("Migrating")
	database.Migrator().CurrentDatabase()

	database.AutoMigrate(&entities.Product{})
	database.AutoMigrate(&entities.ProductImage{})

	database.FirstOrCreate(&entities.Product{
		Name:          "Product Item Test",
		Description:   "Description Test",
		Price:         1245.36,
		DiscountPrice: 1245.36,
		Quantity:      154,
		ProductImage:  []entities.ProductImage{{Source: "https://teste.com.br"}},
	})
}
