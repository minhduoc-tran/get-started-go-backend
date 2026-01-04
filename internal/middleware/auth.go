package middleware

import "github.com/gin-gonic/gin"

// Logger middleware - Ghi log các request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request
		c.Next()
		// Log response
	}
}

// AuthMiddleware - Kiểm tra authorization
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Kiểm tra token
		c.Next()
	}
}
