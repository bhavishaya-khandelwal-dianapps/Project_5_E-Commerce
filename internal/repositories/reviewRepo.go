package repositories

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
)

// 1. Function to submit review
func AddReview(review *models.Review) error {
	if err := config.DB.Create(review).Error; err != nil {
		return err
	}

	return config.DB.Preload("User").Preload("Product").First(review, review.ID).Error
}

// 2. Function to get review
func GetReview(userId, productId uint) (*models.Review, error) {
	var review models.Review
	err := config.DB.Where("user_id = ? AND product_id = ?", userId, productId).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}
