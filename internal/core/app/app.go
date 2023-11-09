package app

import (
	"net/http"
	"time"

	user_controller "khrix/egommerce/internal/modules/user/controller"
	user_repository "khrix/egommerce/internal/modules/user/repository"
	user_service "khrix/egommerce/internal/modules/user/service"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	httpServer := &http.Server{
		Addr:           "127.0.0.1:4000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	userRepo := user_repository.NewUserRepository(nil)
	userService := user_service.NewUserService(userRepo)

	user_controller.NewModule(&router.RouterGroup, userService)

	httpServer.ListenAndServe()
}
