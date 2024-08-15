package infrastructure

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Middleware is a struct that holds the JWTService
type Middleware struct {
	jwtService JWTService
}

// NewMiddleware creates a new Middleware instance
func NewMiddleware(jwtService JWTService) *Middleware {
	return &Middleware{
		jwtService: jwtService,
	}
}

// JWTMiddleware is a Gin middleware for JWT authentication
func (m *Middleware) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Extract token string from the Authorization header
		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		// Validate the token
		token, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims and set them in the context
		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			c.Set("username", (*claims)["username"])
			c.Set("role", (*claims)["role"])
			c.Set("exp", (*claims)["exp"])
			c.Set("user_id", (*claims)["id"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AdminOnly is a middleware that restricts access to admin users
func (m *Middleware) AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
