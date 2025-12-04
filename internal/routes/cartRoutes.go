package routes

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/handlers"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/middleware"
	"github.com/gin-gonic/gin"
)

func CartRoutes(rg *gin.RouterGroup) {
	cart := rg.Group("/cart")
	{
		cart.POST("/", middleware.Auth(), middleware.IsAdminOrUser(), handlers.AddToCart)
		cart.GET("/", middleware.Auth(), middleware.IsAdminOrUser(), handlers.GetCartItems);
	}
}
