package controller

import (
	"net/http"
	"zian-product-api/config"
	"zian-product-api/data/formatter"
	"zian-product-api/data/model"
	"zian-product-api/domain/repository"
	"zian-product-api/middleware"
	"zian-product-api/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService service.ProductService
	repository.ProductRepository
	repository.AuthRepository
	config.Config
}

func NewProductController(productService *service.ProductService, productRepository repository.ProductRepository, authRepository repository.AuthRepository, config config.Config) ProductController {
	return ProductController{ProductService: *productService, ProductRepository: productRepository, AuthRepository: authRepository, Config: config}
}

func (controller *ProductController) Route(app *gin.Engine) {
	api := app.Group("/api//v1")

	api.GET("/product", middleware.AuthenticateJWT(controller.AuthRepository, controller.Config), controller.GetProductList)

}

func (controller *ProductController) GetProductList(c *gin.Context) {
	products, err := controller.ProductService.GetProductList(c.Copy())

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Products Gagal",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Get Products Success",
		Data:    formatter.FormatProductList(products),
	})
	return
}
