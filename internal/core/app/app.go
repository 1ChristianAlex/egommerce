package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	user_auth "khrix/egommerce/internal/core/auth"
	file_upload "khrix/egommerce/internal/libs/file_upload"
	product_controller "khrix/egommerce/internal/modules/catalog/controller"
	product_mapper "khrix/egommerce/internal/modules/catalog/mapper"
	product_repository "khrix/egommerce/internal/modules/catalog/repository"
	product_service "khrix/egommerce/internal/modules/catalog/service"
	user_controller "khrix/egommerce/internal/modules/user/controller"
	user_mapper "khrix/egommerce/internal/modules/user/mapper"
	user_repository "khrix/egommerce/internal/modules/user/repository"
	user_service "khrix/egommerce/internal/modules/user/service"
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

	categoryMapper := product_mapper.NewCategoryMapper()
	productMapper := product_mapper.NewProductMapper(categoryMapper)
	userMapper := user_mapper.NewUserMapper()
	featureItemMapper := product_mapper.NewFeatureItemMapper()
	productFeatureMapper := product_mapper.NewProductFeatureMapper(featureItemMapper)

	userR := user_repository.NewUserRepository(database)
	productR := product_repository.NewProductRepository(database)
	productImageR := product_repository.NewProductImageRepository(database)
	categoryRepository := product_repository.NewCategoryRepository(database)
	productFeatureRepository := product_repository.NewProductFeatureRepository(database)

	passwordS := user_service.NewPasswordService()
	jwtS := user_service.NewJwtService()
	userS := user_service.NewUserService(userR, passwordS, userMapper)
	productImageS := product_service.NewProductImageService(productImageR, productR, file_upload.NewFileUploadManager())

	productFeatureService := product_service.NewProductFeatureService(productFeatureRepository,
		productFeatureMapper,
		featureItemMapper,
		productMapper,
	)
	categoryService := product_service.NewCategoryService(
		categoryRepository,
		categoryMapper,
		productR,
		productMapper,
	)
	productS := product_service.NewProductService(productR, productImageR, productMapper)

	authHelper := user_auth.NewAuthHelper(jwtS)

	apiRouter := router.Group("api", authHelper.JwtMiddleware)

	user_controller.NewAuthController(&router.RouterGroup, userS, jwtS)
	user_controller.NewUserController(apiRouter, userS)
	product_controller.NewProductController(apiRouter, productS, productFeatureService)
	product_controller.NewProductImageController(apiRouter, productS, productImageS)
	product_controller.NewCategoryController(apiRouter, categoryService)

	httpServer.ListenAndServe()
}
