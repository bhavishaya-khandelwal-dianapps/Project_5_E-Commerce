package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/services"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/utils"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Role      string `json:"role,omitempty"`
}

// * 1. Function to register user
func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	if input.Role != "" {
		user.Role = input.Role
	}

	if err := services.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send welcome email
	data := utils.WelcomeUserEmail{
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	go utils.SendEmail("Welcome to Acoustic Store", "internal/views/welcomeUser.html", []string{user.Email}, data)

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"status":  true,
		"user":    user,
	})
}

// * 2. Function to get details of logged in user
func GetUser(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"status":  false,
		})
		return
	}

	// Type assert to User
	// c.Get("user") returns interface{}, which is a generic type in Go.
	// It can hold any value, like a string, int, or struct (models.User in our case).
	// Why we need type assertion?
	// Go is strongly typed, so you can’t use userInterface.ID directly.
	// We need to tell Go: “Hey, I know this interface{} is actually a models.User.”
	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve user",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User profile fetched successfully",
		"status":  true,
		"user":    user,
	})
}

// * 3. Function to update profile
func UpdateProfile(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
			"status":  false,
		})
		return
	}

	// Assert User
	user := userInterface.(models.User)

	var input services.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	updatedUser, err := services.UpdateUserProfile(&user, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"status":  true,
		"user":    updatedUser,
	})
}

// * 4. Function to change password
func ChangePassword(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
			"status":  false,
		})
		return
	}

	user := userInterface.(models.User)

	var input services.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	if err := services.ChangePassword(&user, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed successfully",
		"status":  true,
	})
}

// * 5. Function to get all users
func GetAllUsers(c *gin.Context) {
	// Parse query params
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")
	role := c.Query("role")
	sortBy := c.DefaultQuery("sortBy", "created_at")
	sortOrder := c.DefaultQuery("sortOrder", "desc")

	users, totalUsers, err := services.GetAllUsers(pageStr, limitStr, search, role, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	// Convert page and limit to int
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"status":  true,
		"users":   users,
		"pagination": gin.H{
			"total":      totalUsers,
			"page":       page,
			"limit":      limit,
			"totalPages": int(math.Ceil(float64(totalUsers) / float64(limit))),
		},
	})
}

// * 6. Function to delete user by id
func DeleteUser(c *gin.Context) {
	idParams := c.Param("id")

	// Convert id to uint
	id, err := strconv.ParseUint(idParams, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
			"status":  false,
		})
		return
	}

	// Call service
	err = services.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"status":  true,
	})
}
