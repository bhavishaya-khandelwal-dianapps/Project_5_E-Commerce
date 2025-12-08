package services

import (
	"errors"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/repositories"
)

// 1. Function to add review
type SubmitReviewInput struct {
	ProductId uint    `json:"productId" binding:"required"`
	Rating    int     `json:"rating" binding:"required,gte=0,lte=5"`
	Comment   *string `json:"comment" binding:"omitempty"`
}

func AddReview(userId uint, input *SubmitReviewInput) (*models.Review, error) {
	// 1. Check if product exist or not
	product, err := repositories.GetProduct(uint(input.ProductId))
	if err != nil {
		return nil, err
	}

	// 2. Check is already reviewed
	existingReview, err := repositories.GetReview(userId, product.ID)
	if err == nil && existingReview != nil {
		return nil, errors.New("you already reviewed this product")
	}

	var comment = ""
	if input.Comment != nil {
		comment = *input.Comment
	}

	review := models.Review{
		UserId:    userId,
		ProductId: product.ID,
		Rating:    input.Rating,
		Comment:   comment,
	}

	err = repositories.AddReview(&review)
	if err != nil {
		return nil, err
	}

	return &review, nil
}
