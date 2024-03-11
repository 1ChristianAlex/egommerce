package migrations

import (
	"fmt"
	"time"

	service "khrix/egommerce/internal/application/user/services"
	"khrix/egommerce/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	fmt.Println("Migrating")
	database.Migrator().CurrentDatabase()

	database.AutoMigrate(&models.User{})
	password, _ := service.NewPasswordService().HashPassword("123456789")

	database.FirstOrCreate(&models.User{
		Username: "admin",
		Password: password,
		Name:     "Christian Alexsander",
		Email:    "christian.alexsander@outlook.com",
		Birthday: time.Date(1999, time.June, 13, 0, 0, 0, 0, time.UTC),
	})

	ProductMigration(database)
}
