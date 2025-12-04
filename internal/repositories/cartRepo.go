package repositories

import (
	"errors"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"gorm.io/gorm"
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

func DeleteCartItem(cart *models.Cart) error {
	return config.DB.Delete(cart).Error
}

func RemoveProductFromCart(userId, productId uint) error {
	return config.DB.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&models.Cart{}).Error
}

type UpdateCartByUserIdInput struct {
	ProductId uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"min=0"`
}

func UpdateCartByUserId(userId uint, input UpdateCartByUserIdInput) (*models.Cart, error) {
	var cart models.Cart

	// 1. Find the cart item
	err := config.DB.Preload("User").Preload("Product").Where("user_id = ? AND product_id = ?", userId, input.ProductId).First(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("cart item not found")
		}
		return nil, err
	}

	// 2. Update quantity if possible
	if cart.Product.Stock < input.Quantity {
		return nil, errors.New("requested quantity exceeds available stock")
	}
	cart.Quantity = input.Quantity

	if err := config.DB.Save(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func ClearCart(userId uint) error {
	var cartItems []models.Cart

	err := config.DB.Where("user_id = ?", userId).Find(&cartItems).Error
	if err != nil {
		return err
	}

	// Check if cart is already empty
	if len(cartItems) == 0 {
		return errors.New("cart is already empty")
	}

	err = config.DB.Where("user_id = ?", userId).Delete(&models.Cart{}).Error
	if err != nil {
		return err
	}

	return nil
}
