package middleware

import (
	"errors"
	"strings"
	"zian-product-api/config"
	"zian-product-api/data/model"
	"zian-product-api/domain/repository"
	"zian-product-api/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthenticateJWT(repository repository.AuthRepository, config config.Config) gin.HandlerFunc {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.AbortWithStatusJSON(400, model.GeneralResponse{
				Code:    400,
				Message: "Bad Request",
				Data:    "Missing or malformed JWT",
			})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Invalid Token")
			}

			return []byte(jwtSecret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(400, model.GeneralResponse{
				Code:    400,
				Message: "Bad Requests",
				Data:    "Missing or malformed JWT",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(401, model.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    "Invalid or expired JWT",
			})
			return
		}
		memberID := int(claims["user_id"].(float64))
		member, err := service.NewAuthServiceImpl(repository).GetUserByID(c.Copy(), memberID)

		if err != nil {
			c.AbortWithStatusJSON(401, model.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    "Invalid or expired JWT",
			})
			return

		}
		c.Set("currentUser", member)

	}
}
