package routes

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/handlers"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/middleware"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(rg *gin.RouterGroup) {
	payment := rg.Group("/payment")
	{
		payment.POST("/create", middleware.Auth(), middleware.IsUser(), handlers.CreatePayment)
		payment.POST("/verify", middleware.Auth(), middleware.IsUser(), handlers.VerifyPayment)
	}
}
