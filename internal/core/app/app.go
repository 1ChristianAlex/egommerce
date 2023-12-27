package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	file_manager "khrix/egommerce/internal/libs/file_manager"
	file_upload "khrix/egommerce/internal/libs/file_upload"
	auth_controller "khrix/egommerce/internal/modules/auth/controller"
	auth_helper "khrix/egommerce/internal/modules/auth/helper"
	auth_service "khrix/egommerce/internal/modules/auth/service"
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

	featureItemMapper := product_mapper.NewFeatureItemMapper()
	productFeatureMapper := product_mapper.NewProductFeatureMapper(featureItemMapper)
	categoryMapper := product_mapper.NewCategoryMapper()
	productMapper := product_mapper.NewProductMapper(categoryMapper, featureItemMapper)
	userMapper := user_mapper.NewUserMapper()

	userR := user_repository.NewUserRepository(database)
	productR := product_repository.NewProductRepository(database)
	productImageR := product_repository.NewProductImageRepository(database)
	categoryRepository := product_repository.NewCategoryRepository(database)
	productFeatureRepository := product_repository.NewProductFeatureRepository(database)
	productSearchRepository := product_repository.NewSearchRepository(database)

	passwordS := user_service.NewPasswordService()
	jwtS := auth_service.NewJwtService()
	userS := user_service.NewUserService(userR, passwordS, userMapper)
	fileManager := file_manager.NewFileManager()

	awsFileUpload := file_upload.NewAwsBuckerManager(
		fileManager,
		file_upload.S3Data{
			Bucket:       os.Getenv("BUCKER_NAME"),
			Key:          os.Getenv("BUCKER_KEY_IMAGES"),
			BuckerRegion: os.Getenv("BUCKER_REGION"),
		},
	)

	productImageS := product_service.NewProductImageService(
		productImageR,
		productR,
		awsFileUpload,
	)

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
	productS := product_service.NewProductService(productR, productSearchRepository, productImageR, productMapper)

	authHelper := auth_helper.NewAuthHelper(jwtS)

	apiRouter := router.Group("api", authHelper.JwtMiddleware)

	auth_controller.NewAuthController(&router.RouterGroup, userS, jwtS)
	user_controller.NewUserController(apiRouter, userS)
	product_controller.NewProductController(apiRouter, productS, productFeatureService, authHelper)
	product_controller.NewProductImageController(apiRouter, productS, productImageS)
	product_controller.NewCategoryController(apiRouter, categoryService)

	httpServer.ListenAndServe()
}
