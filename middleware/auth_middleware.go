package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

// AuthMiddleware memvalidasi Backend JWT Token di setiap request
// Dipasang di route group yang memerlukan autentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":    false,
				"message":    "Authorization header tidak ditemukan",
				"error_code": "MISSING_TOKEN",
			})
			return
		}
		// 2. Validasi format "Bearer <token>" 
		parts := strings.SplitN(authHeader, " ", 2) 
		if len(parts) != 2 || parts[0] != "Bearer" { 
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ 
		"success":    false,
		"message": "Format token salah. Gunakan: Bearer <token>", 
		"error_code": "INVALID_TOKEN_FORMAT",
		})
		return 
	}
	tokenString := parts[1]
	
}
