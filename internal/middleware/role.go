package middleware

import (
	"net/http"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/gin-gonic/gin"
)

// isUser middleware - allows only users with role "CUSTOMER"
func IsUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInterface, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			c.Abort()
			return
		}

		user, ok := userInterface.(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user from context"})
			c.Abort()
			return
		}

		if user.Role != "CUSTOMER" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User role required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// isAdmin middleware - allows only users with role "ADMIN"
func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInterface, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			c.Abort()
			return
		}

		user, ok := userInterface.(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user from context"})
			c.Abort()
			return
		}

		if user.Role != "ADMIN" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Admin role required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// isAdminOrUser middleware - allows both "ADMIN" and "CUSTOMER"
func IsAdminOrUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInterface, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			c.Abort()
			return
		}

		user, ok := userInterface.(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user from context"})
			c.Abort()
			return
		}

		if user.Role != "CUSTOMER" && user.Role != "ADMIN" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
