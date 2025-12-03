package routes

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/handlers"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	{
		user.GET("/me", middleware.Auth(), middleware.IsUser(), handlers.GetUser)
		user.PUT("/me", middleware.Auth(), middleware.IsUser(), handlers.UpdateProfile)
		user.PUT("/change-password", middleware.Auth(), middleware.IsUser(), handlers.ChangePassword)
		user.GET("/all", middleware.Auth(), middleware.IsAdmin(), handlers.GetAllUsers)
	}
}
