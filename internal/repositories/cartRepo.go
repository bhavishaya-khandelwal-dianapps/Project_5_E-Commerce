package repositories

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
)

func GetCartItem(userId, productId uint) (*models.Cart, error) {
	var cart models.Cart

	err := config.DB.
		Where("user_id = ? AND product_id = ?", userId, productId).
		Preload("User").    // <-- preload the User data
		Preload("Product"). // <-- preload the Product data
		First(&cart).Error

	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func AddToCart(cart *models.Cart) error {
	return config.DB.Create(cart).Error
}

func UpdateCart(cart *models.Cart) error {
	return config.DB.Save(cart).Error
}

func GetCartItems(userId uint) ([]models.Cart, error) {
	var cartItems []models.Cart
	err := config.DB.Preload("Product").Preload("User").Where("user_id = ?", userId).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	return cartItems, err
}
