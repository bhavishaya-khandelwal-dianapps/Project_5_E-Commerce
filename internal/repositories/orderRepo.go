package repositories

import (
	"errors"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"gorm.io/gorm"
)

// 1. Function to create order
func CreateOrder(order *models.Order) error {
	return config.DB.Preload("User").Create(order).Error
}

// 2. Function to get order by id
func GetOrderById(orderId uint) (*models.Order, error) {
	var order models.Order
	err := config.DB.First(&order, orderId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return &order, nil
}

// 3. Function to update order
func UpdateOrder(order *models.Order) error {
	return config.DB.Save(order).Error
}
