package routes

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/handlers"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(rg *gin.RouterGroup) {
	order := rg.Group("/order")
	{
		order.POST("/create", middleware.Auth(), middleware.IsUser(), handlers.PlaceOrder)
	}
}
