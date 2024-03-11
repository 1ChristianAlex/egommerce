package core

import (
	"os"

	"khrix/egommerce/internal/infrastructure"
	"khrix/egommerce/internal/infrastructure/database"
	"khrix/egommerce/internal/infrastructure/database/migrations"
	"khrix/egommerce/internal/infrastructure/server"

	auth_controller "khrix/egommerce/internal/adapter/auth/controller"
	product_controller "khrix/egommerce/internal/adapter/catalog/controller"
	product_mapper "khrix/egommerce/internal/adapter/catalog/mapper"
	product_repository "khrix/egommerce/internal/adapter/catalog/repository"
	user_controller "khrix/egommerce/internal/adapter/user/controller"
	user_mapper "khrix/egommerce/internal/adapter/user/mapper"
	user_repository "khrix/egommerce/internal/adapter/user/repository"

	file_upload "khrix/egommerce/internal/adapter/libs/file_upload"
	product_service "khrix/egommerce/internal/application/catalog/services"
	file_manager "khrix/egommerce/internal/application/libs/file_manager"
	user_service "khrix/egommerce/internal/application/user/services"
	auth_service "khrix/egommerce/internal/domain/auth/service"

	auth_helper "khrix/egommerce/internal/infrastructure/server/auth"
)

func StartApp() {
	infrastructure.LoadEnvFile()
	dbPool := database.ConnectDatabase()
	migrations.Migrate(dbPool)

	featureItemMapper := product_mapper.NewFeatureItemMapper()
	productFeatureMapper := product_mapper.NewProductFeatureMapper(featureItemMapper)
	categoryMapper := product_mapper.NewCategoryMapper()
	productMapper := product_mapper.NewProductMapper(categoryMapper, featureItemMapper)
	userMapper := user_mapper.NewUserMapper()

	userR := user_repository.NewUserRepository(dbPool)
	productR := product_repository.NewProductRepository(dbPool)
	productImageR := product_repository.NewProductImageRepository(dbPool)
	categoryRepository := product_repository.NewCategoryRepository(dbPool)
	productFeatureRepository := product_repository.NewProductFeatureRepository(dbPool)
	productSearchRepository := product_repository.NewSearchRepository(dbPool)

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

	server.StartServer(
		auth_controller.NewAuthController(userS, jwtS),
		user_controller.NewUserController(userS),
		product_controller.NewProductController(productS, productFeatureService, authHelper),
		product_controller.NewProductImageController(productS, productImageS),
		product_controller.NewCategoryController(categoryService),
		auth_helper.NewAuthHelper(jwtS),
	)
}
