package common

import (
	"strconv"
	"time"
	"zian-product-api/common/exception"
	"zian-product-api/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, userID int, config config.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicIfNeeded(err)
	claims := jwt.MapClaims{
		"username": username,
		"user_id":  userID,
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicIfNeeded(err)

	return tokenSigned
}
