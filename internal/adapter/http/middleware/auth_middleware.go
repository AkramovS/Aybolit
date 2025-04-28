package middleware

import (
	"Aybolit/pkg/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			log.Println("Authorization header format must be Bearer {token}")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		log.Println("[AuthMiddleware] Token:", tokenString)

		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			log.Println("[AuthMiddleware] Token:", tokenString)
			return
		}

		c.Set("login", claims.Login)
		c.Set("role", claims.Role)
		c.Next()
	}
}

/*
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		log.Println(tokenString)
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("login", claims.Login)
		c.Set("role", claims.Role)
		c.Next()
	}
}
*/
