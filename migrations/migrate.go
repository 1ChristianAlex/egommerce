package migrations

import (
	"fmt"
	"time"

	"khrix/egommerce/internal/modules/user/repository/entities"
	"khrix/egommerce/internal/modules/user/service"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	fmt.Println("Migrating")
	database.Migrator().CurrentDatabase()

	database.AutoMigrate(&entities.User{})
	password, _ := service.NewPasswordService().HashPassword("123456789")

	database.FirstOrCreate(&entities.User{
		Username: "admin",
		Password: password,
		Name:     "Christian Alexsander",
		Email:    "christian.alexsander@outlook.com",
		Birthday: time.Date(1999, time.June, 13, 0, 0, 0, 0, time.UTC),
	})

	ProductMigration(database)
	CategoryMigration(database)
}
