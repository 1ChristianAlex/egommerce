package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	user_controller "khrix/egommerce/internal/modules/user/controller"
	user_repository "khrix/egommerce/internal/modules/user/repository"
	user_service "khrix/egommerce/internal/modules/user/service"
	"khrix/egommerce/migrations"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	LoadEnvFile()

	database := ConnectDatabase()
	migrations.Migrate(database)

	router := gin.Default()

	address := fmt.Sprintf("%s:%s", os.Getenv("API_URL"), os.Getenv("API_PORT"))

	httpServer := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	userRepo := user_repository.NewUserRepository(database)
	userService := user_service.NewUserService(userRepo)

	user_controller.NewModule(&router.RouterGroup, userService)

	httpServer.ListenAndServe()
}
