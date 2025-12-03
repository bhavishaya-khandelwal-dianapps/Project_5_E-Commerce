package repositories

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
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

// 3. Function to get product by id
func GetProduct(id uint) (*models.Product, error) {
	var product models.Product
	err := config.DB.Find(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
