package repositories

import (
	"errors"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"gorm.io/gorm"
)

// 1. Function to create product
func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

// 2. Function to get products
type ProductQueryParams struct {
	Search    string
	PriceMin  *float64
	PriceMax  *float64
	StockMin  *int
	StockMax  *int
	SortBy    string
	SortOrder string
	Page      int
	Limit     int
}

func GetAllProducts(params ProductQueryParams) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := config.DB.Model(&models.Product{})

	// Search by name
	if params.Search != "" {
		query = query.Where("name ILIKE ?", "%"+params.Search+"%")
	}

	// Filter by price
	if params.PriceMin != nil {
		query = query.Where("price >= ?", params.PriceMin)
	}

	if params.PriceMax != nil {
		query = query.Where("price <= ?", params.PriceMax)
	}

	// Filter by stock
	if params.StockMin != nil {
		query = query.Where("stock >= ?", params.StockMin)
	}

	if params.StockMax != nil {
		query = query.Where("stock <= ?", params.StockMax)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (params.Page - 1) * params.Limit
	query = query.Order(params.SortBy + " " + params.SortOrder).Offset(offset).Limit(params.Limit)

	// Execute Query
	if err := query.Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// Define a custom error
var ErrProductNotFound = errors.New("product not found")

// 3. Function to get product by id
func GetProduct(id uint) (*models.Product, error) {
	var product models.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

// 4. Function to update product
func UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

// 5. Function to delete product
func DeleteProduct(product *models.Product) error {
	return config.DB.Delete(product).Error
}

// 6. Function to decrease stock
func DecreaseProductStock(productId uint, qty int) error {
	var product models.Product

	// Fetch product
	if err := config.DB.First(&product, productId).Error; err != nil {
		return err
	}

	// Check stock
	if product.Stock < qty {
		return errors.New("insufficient stock")
	}

	// Update stock
	return config.DB.Model(&product).Update("stock", product.Stock-qty).Error
}
