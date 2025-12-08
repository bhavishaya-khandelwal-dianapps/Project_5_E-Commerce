package services

import (
	"errors"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/repositories"
	"gorm.io/gorm"
)

// 1. Function to create order
func CreateOrder(userId uint) (*models.Order, error) {
	// 1. Fetch cart items
	cartItems, err := repositories.GetCartItems(userId)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New("oops, your cart is empty")
	}

	// 2. Calculate total price
	var total float64 = 0
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Product.Price
	}

	// 3. Create order
	order := &models.Order{
		UserId:     userId,
		TotalPrice: total,
		Status:     "PENDING",
	}

	err = repositories.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	// 4. Create order items + reduce stock
	for _, item := range cartItems {
		orderItem := models.OrderItem{
			OrderId:   order.ID,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		}

		err := repositories.CreateOrderItem(&orderItem)
		if err != nil {
			return nil, err
		}

		// Reduce Stock
		err = repositories.DecreaseProductStock(uint(item.ProductId), item.Quantity)
		if err != nil {
			return nil, err
		}
	}

	// 5. Clear cart
	err = repositories.ClearCart(userId)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func GetOrderById(orderId uint) (*models.Order, error) {
	return repositories.GetOrderById(orderId)
}

// 2. Function to cancel order
func CancelOrder(orderId, userId uint) error {
	order, err := repositories.GetOrderById(orderId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("order not found")
		}
		return err
	}

	if order.UserId != userId {
		return errors.New("oops, this is not your order")
	}

	if order.Status != "PENDING" {
		return errors.New("cannot cancel this order")
	}

	order.Status = "CANCELLED"
	err = repositories.UpdateOrder(order)
	if err != nil {
		return err
	}

	return nil
}
