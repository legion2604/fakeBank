package middleware

import (
	"fakeBank/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No access token"})
			c.Abort()
			return
		}
		accessSecret := []byte(config.GetEnv("JWT_SECRET_ACCESS"))

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return accessSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userIdFloat, ok := claims["userId"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("userId", int(userIdFloat))
		c.Next()
	}
}

func GetUserIDFromContext(c *gin.Context) int {
	id := c.GetInt("userId") // достаём значение из контекста
	return id
}
