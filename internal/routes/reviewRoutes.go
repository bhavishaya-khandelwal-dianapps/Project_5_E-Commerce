package routes

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/handlers"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/middleware"
	"github.com/gin-gonic/gin"
)

func ReviewRoutes(rg *gin.RouterGroup) {
	review := rg.Group("/review")
	{
		review.POST("/add", middleware.Auth(), middleware.IsUser(), handlers.SubmitReview)
	}
}
