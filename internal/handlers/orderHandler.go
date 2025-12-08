package handlers

import (
	"net/http"
	"strconv"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/services"
	"github.com/gin-gonic/gin"
)

// 1. Function to create order
func PlaceOrder(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
			"status":  false,
		})
		return
	}

	user := userInterface.(models.User)
	userId := user.ID

	// Call Service
	order, err := services.CreateOrder(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"status":  true,
		"order":   order,
	})
}

// 2. Function to cancel order
func CancelOrder(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
			"status":  false,
		})
		return
	}

	user := userInterface.(models.User)
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order id",
			"status":  false,
		})
		return
	}

	if err := services.CancelOrder(uint(id), uint(user.ID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order cancelled successfully",
		"status":  true,
	})
}
