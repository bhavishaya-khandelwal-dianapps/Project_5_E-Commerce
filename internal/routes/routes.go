package routes

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/handlers"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		// Auth Routes
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		// User Routes
		api.GET("/user/me", middleware.Auth(), middleware.IsUser(), handlers.GetUser)
		api.PUT("/user/me", middleware.Auth(), middleware.IsUser(), handlers.UpdateProfile)
		api.PUT("/user/change-password", middleware.Auth(), middleware.IsUser(), handlers.ChangePassword)
		api.GET("/users", middleware.Auth(), middleware.IsAdmin(), handlers.GetAllUsers)
	}
}
