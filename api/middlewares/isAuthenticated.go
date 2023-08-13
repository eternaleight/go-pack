package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "権限がありません。"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, bearerSchema)
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("YOUR_SECRET_KEY"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["id"].(string) // or whichever type you set the user ID to be
			c.Set("userId", userId)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "権限がありません。"})
			c.Abort()
			return
		}
	}
}
