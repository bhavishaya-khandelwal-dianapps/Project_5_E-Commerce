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
	Stock       int     `json:"stock" binding:"required"`
	ImageURL    string  `json:"imageUrl"`
}

func CreateProduct(input *CreateProductInput) (*models.Product, error) {
	product := &models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		ImageURL:    input.ImageURL,
	}

	if err := repositories.CreateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

// * 2. Function to get all products
func GetAllProducts(params repositories.ProductQueryParams) ([]models.Product, int64, error) {
	return repositories.GetAllProducts(params)
}

// * 3. Function to get product by id
func GetProduct(id uint) (*models.Product, error) {
	product, err := repositories.GetProduct(id)
	if err != nil {
		if err == repositories.ErrProductNotFound {
			return nil, repositories.ErrProductNotFound
		}
		return nil, err
	}

	return product, nil
}

// * 4. Function to update product
type UpdateProductInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *int     `json:"stock"`
	ImageURL    string   `json:"imageUrl"`
}

func UpdateProduct(id uint, input UpdateProductInput) (*models.Product, error) {
	product, err := GetProduct(id)
	if err != nil {
		// Check if error is record not found
		if err == repositories.ErrProductNotFound {
			return nil, repositories.ErrProductNotFound
		}
		return nil, err
	}

	if input.Name != "" {
		product.Name = input.Name
	}

	if input.Description != "" {
		product.Description = input.Description
	}

	if input.Price != nil {
		product.Price = *input.Price
	}

	if input.Stock != nil {
		product.Stock = *input.Stock
	}

	if input.ImageURL != "" {
		product.ImageURL = input.ImageURL
	}

	if err := repositories.UpdateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

//* 5. Function to delete product 
func DeleteProduct(id uint) error {
	product, err := GetProduct(id)
	if err != nil {
		if err == repositories.ErrProductNotFound {
			return repositories.ErrProductNotFound
		}
		return err 
	}

	return repositories.DeleteProduct(product);
} 