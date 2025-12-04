package handlers

import (
	"net/http"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/services"
	"github.com/gin-gonic/gin"
)

// 1. Function to add product into the cart
func AddToCart(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
			"status":  false,
		})
		return
	}

	user := userInterface.(models.User)

	var input services.AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	cartItem, err := services.AddToCart(user.ID, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Product added to cart successfully",
		"status":   true,
		"cartItem": cartItem,
	})
}

// 2. Function to get cart items
func GetCartItems(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"status":  false,
		})
		return
	}

	user := userInterface.(models.User)
	userId := user.ID

	cartItems, err := services.GetCartItems(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch cart items",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Cart items fetched successfully",
		"status":    true,
		"cartItems": cartItems,
	})
}
