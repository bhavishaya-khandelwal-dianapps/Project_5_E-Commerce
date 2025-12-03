package handlers

import (
	"net/http"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/services"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	// Validate incoming JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		// Return user-friendly error
		var message string

		// Check for missing fields
		if req.Email == "" {
			message = "Email is required and must be valid"
		} else if req.Password == "" {
			message = "Password is required"
		} else {
			message = "Invalid request"
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
			"status":  false,
		})
		return
	}

	// Call login service
	token, user, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"status":  true,
		"token":   token,
		"user":    user,
	})
}
