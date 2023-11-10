package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"khrix/egommerce/internal/modules/user/controller"
	"khrix/egommerce/internal/modules/user/repository"
	"khrix/egommerce/internal/modules/user/service"
	"khrix/egommerce/migrations"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	LoadEnvFile()

	database := ConnectDatabase()
	migrations.Migrate(database)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	address := fmt.Sprintf("%s:%s", os.Getenv("API_URL"), os.Getenv("API_PORT"))

	httpServer := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	passwordService := service.NewPasswordService()
	jwtService := service.NewJwtService()
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo, passwordService)

	controller.NewModule(&router.RouterGroup, userService, jwtService)

	httpServer.ListenAndServe()
}
