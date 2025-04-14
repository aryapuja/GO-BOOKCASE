package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the "Authorization" header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Split "Bearer <token>" and get the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		// Parse and validate JWT token
		claims := jwt.MapClaims{}
		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "SECRET_KEY is not set in the environment variables"})
			c.Abort()
			return
		}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// If token is valid, set claims to context and proceed with the request
		c.Set("claims", claims)
		c.Next()
	}
}
