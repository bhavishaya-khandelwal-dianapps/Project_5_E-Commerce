package repositories

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
)

func CreateOrderItem(orderItem *models.OrderItem) error {
	return config.DB.Create(orderItem).Error
}
