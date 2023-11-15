package controller

import (
	"net/http"
	"zian-product-api/common"
	"zian-product-api/common/exception"
	"zian-product-api/config"
	"zian-product-api/data/model"
	"zian-product-api/domain/repository"
	"zian-product-api/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthService
	repository.AuthRepository
	config.Config
}

func NewAuthController(authService *service.AuthService, authRepository repository.AuthRepository, config config.Config) AuthController {
	return AuthController{AuthService: *authService, AuthRepository: authRepository, Config: config}
}

func (controller *AuthController) Route(app *gin.Engine) {
	api := app.Group("/api//v1")

	api.POST("/register", controller.Register)
	api.POST("/login", controller.Login)
}

func (controller *AuthController) Register(c *gin.Context) {
	var input model.RegisterBody

	errReq := c.ShouldBindJSON(&input)
	exception.PanicIfNeeded(errReq)

	if errReq != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Register Gagal",
			Data:    nil,
		})
		return
	}

	emailAvailable, _ := controller.AuthService.CheckEmailAvailable(c.Copy(), input.Email)

	if emailAvailable {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Email sudah terdaftar silahkan gunakan email yang lain",
			Data:    nil,
		})
		return
	}

	newUser := controller.AuthService.Register(c.Copy(), input)

	jwtToken := common.GenerateToken(newUser.Name, newUser.ID, controller.Config)

	resultWithToken := map[string]interface{}{
		"username": newUser.Name,
		"token":    jwtToken,
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Register Berhasil",
		Data:    resultWithToken,
	})
	return
}

func (controller *AuthController) Login(c *gin.Context) {
	var input model.LoginBody

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Login Gagal",
			Data:    nil,
		})
		return
	}

	_, err = controller.AuthService.CheckEmailOrPasswordValid(c.Copy(), input)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Email atau password salah",
			Data:    nil,
		})
		return
	}

	response, errs := controller.AuthService.Login(c.Copy(), input)
	if errs != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Gagal",
			Data:    nil,
		})
		return
	}
	tokenJwtResult := common.GenerateToken(response.Name, response.ID, controller.Config)
	resultWithToken := map[string]interface{}{
		"user_id":      response.ID,
		"name":         response.Name,
		"email":        response.Email,
		"access_token": tokenJwtResult,
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Login Berhasil",
		Data:    resultWithToken,
	})
	return
}
