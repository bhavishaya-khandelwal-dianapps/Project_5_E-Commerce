package services

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/repositories"
)

// * 1. Function to create product
type CreateProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int      `json:"stock" binding:"required"`
	ImageURL    string   `json:"imageUrl"`
}

func CreateProduct(input *CreateProductInput) (*models.Product, error) {
	product := &models.Product{
		Name: input.Name, 
		Description: input.Description,
		Price: input.Price,
		Stock: input.Stock,
		ImageURL: input.ImageURL,
	}

	if err := repositories.CreateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

//* 2. Function to get all products 
func GetAllProducts(params repositories.ProductQueryParams) ([]models.Product, int64, error) {
	return repositories.GetAllProducts(params)
}

//* 3. Function to get product by id 
func GetProduct(id uint) (*models.Product, error) {
	return repositories.GetProduct(id)
}