package services

import (
	"errors"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/repositories"
	"gorm.io/gorm"
)

// 1. Function to add product in cart
type AddToCartInput struct {
	ProductId uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

func AddToCart(userId uint, input *AddToCartInput) (*models.Cart, error) {
	// 1. Check if product exists
	product, err := repositories.GetProduct(input.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// 2. Check if cart item already exists
	cartItem, err := repositories.GetCartItem(userId, input.ProductId)
	if err == nil {
		// Cart exists → calculate new total quantity
		newQuantity := cartItem.Quantity + input.Quantity
		if newQuantity > product.Stock {
			return nil, errors.New("requested quantity exceeds available stock")
		}
		cartItem.Quantity = newQuantity

		if err := repositories.UpdateCart(cartItem); err != nil {
			return nil, err
		}
		return cartItem, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Any other DB error
		return nil, err
	}

	// 3. Cart item does not exist → check stock
	if input.Quantity > product.Stock {
		return nil, errors.New("requested quantity exceeds available stock")
	}

	newCart := &models.Cart{
		UserId:    userId,
		ProductId: product.ID,
		Quantity:  input.Quantity,
	}

	if err := repositories.AddToCart(newCart); err != nil {
		return nil, err
	}

	return newCart, nil
}

// 2. Function to get cart items
func GetCartItems(userId uint) ([]models.Cart, error) {
	return repositories.GetCartItems(userId)
}

// 3. Function to update cart
func UpdateCartByUserId(userId uint, input repositories.UpdateCartByUserIdInput) (*models.Cart, error) {
	return repositories.UpdateCartByUserId(userId, input)
}

// 4. Function to remove item from cart
func DeleteProductFromCart(userId, productId uint) error {
	return repositories.RemoveProductFromCart(userId, productId)
}

// 5. Function to clear cart
func ClearCart(userId uint) error {
	return repositories.ClearCart(userId)
}
