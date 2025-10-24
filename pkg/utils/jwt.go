package utils

import (
	"fakeBank/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 12 hours
	})
	return token.SignedString([]byte(config.GetEnv("JWT_SECRET_ACCESS")))
}
