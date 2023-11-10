package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	userAuth "khrix/egommerce/internal/core/auth"
	productController "khrix/egommerce/internal/modules/product/controller"
	productRepository "khrix/egommerce/internal/modules/product/repository"
	productService "khrix/egommerce/internal/modules/product/service"
	userController "khrix/egommerce/internal/modules/user/controller"
	userRepository "khrix/egommerce/internal/modules/user/repository"
	userService "khrix/egommerce/internal/modules/user/service"
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

	userR := userRepository.NewUserRepository(database)
	productR := productRepository.NewProductRepository(database)
	productImageR := productRepository.NewProductImageRepository(database)

	passwordS := userService.NewPasswordService()
	jwtS := userService.NewJwtService()
	userS := userService.NewUserService(userR, passwordS)

	productS := productService.NewProductService(productR, productImageR)

	authHelper := userAuth.NewAuthHelper(jwtS)

	apiRouter := router.Group("api", authHelper.JwtMiddleware)

	userController.NewAuthModule(&router.RouterGroup, userS, jwtS)
	userController.NewUserModule(apiRouter, userS)
	productController.NewModule(apiRouter, productS)

	httpServer.ListenAndServe()
}
