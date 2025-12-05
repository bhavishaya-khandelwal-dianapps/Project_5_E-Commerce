package repositories

import (
	"errors"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
)

// 1. Function to save payment
func CreatePayment(payment *models.Payment) error {
	return config.DB.Create(payment).Error
}

// 2. Function to get payment by razorpay order id
func GetPaymentByRazorpayOrderId(orderId string) (*models.Payment, error) {
	var payment models.Payment
	err := config.DB.Where("razorpay_order_id = ?", orderId).First(&payment).Error
	if err != nil {
		return nil, errors.New("payment record not found")
	}
	return &payment, nil
}

// 3. Function to update payment
func UpdatePayment(payment *models.Payment) error {
	return config.DB.Save(payment).Error
}
