package migrations

import (
	"fmt"

	"khrix/egommerce/internal/modules/categories/repository/entities"

	"gorm.io/gorm"
)

func CategoryMigration(database *gorm.DB) {
	fmt.Println("Migrating")
	database.Migrator().CurrentDatabase()

	database.AutoMigrate(&entities.Category{})

	database.FirstOrCreate(&entities.Category{
		Name:     "Test Category",
		Category: []entities.Category{{Name: "Test Sub Category"}},
	})
}
