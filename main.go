package main

import (
	"zian-product-api/common/exception"
	"zian-product-api/config"
	"zian-product-api/controller"
	"zian-product-api/domain/repository"
	"zian-product-api/middleware"
	"zian-product-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	configuration := config.New()
	db := config.NewDatabase(configuration)

	// repository
	authRepository := repository.NewAuthRepositoryImpl(db)
	productRepository := repository.NewProductRepositoryImpl(db)

	// service
	authService := service.NewAuthServiceImpl(authRepository)
	productService := service.NewProductServiceImpl(productRepository)

	// controller
	authController := controller.NewAuthController(&authService, authRepository, configuration)
	productController := controller.NewProductController(&productService, productRepository, authRepository, configuration)

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Static("/storage", "./storage")
	app.Use(gin.CustomRecovery(exception.ErrorHandler))
	app.Use(middleware.CORSMiddleware())

	// setup routing
	authController.Route(app)
	productController.Route(app)

	// start app
	err := app.Run(":9090")
	exception.PanicIfNeeded(err)
}
