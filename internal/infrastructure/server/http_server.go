package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	auth_helper "khrix/egommerce/internal/infrastructure/server/auth"
	auth_port "khrix/egommerce/internal/port/auth"
	catalog_port "khrix/egommerce/internal/port/catalog"
	user_port "khrix/egommerce/internal/port/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func bindHttpAuth(router *gin.RouterGroup, controller auth_port.AuthController[gin.Context]) {
	router.POST("/login", controller.DoLogin)
}

func bindHttpUser(router *gin.RouterGroup, controller user_port.UserController[gin.Context]) {
	router.POST("/create", controller.CreateNewUser)
}

func bindHttpProduct(router *gin.RouterGroup, controller catalog_port.ProductController[gin.Context]) {
	router.POST("/", controller.CreateNewProductItem)
	router.GET("/", controller.GetListProducts)
}

func bindHttpProductImageController(router *gin.RouterGroup, controller catalog_port.ProductImageController[gin.Context]) {
	router.PUT("/product-image/:productId", controller.UploadImage)
}

func bindHttpCategoryController(router *gin.RouterGroup, controller catalog_port.CategoryController[gin.Context]) {
	router.POST("/create", controller.CreateNewCategory)
	router.POST("/set-product", controller.SetProductCategory)
	router.GET("/", controller.ListAllCategories)
	router.GET("/:categoryId", controller.ListProductsFromCategory)
}

func StartServer(
	authController auth_port.AuthController[gin.Context],
	userController user_port.UserController[gin.Context],
	productController catalog_port.ProductController[gin.Context],
	productImageController catalog_port.ProductImageController[gin.Context],
	categoryController catalog_port.CategoryController[gin.Context],
	authHelper *auth_helper.AuthHelper,
) {
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

	apiRouter := router.Group("api", authHelper.JwtMiddleware)

	bindHttpAuth(router.Group(""), authController)
	bindHttpUser(apiRouter.Group("user"), userController)
	bindHttpProduct(apiRouter.Group("product"), productController)
	bindHttpProductImageController(apiRouter.Group("product"), productImageController)

	bindHttpCategoryController(apiRouter.Group("/category"), categoryController)

	httpServer.ListenAndServe()
}
