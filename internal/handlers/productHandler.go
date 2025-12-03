package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/repositories"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/services"
	"github.com/gin-gonic/gin"
)

// 1. Function to create product
func CreateProduct(c *gin.Context) {
	var input services.CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	product, err := services.CreateProduct(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
		"status":  true,
		"product": product,
	})
}

// 2. Function to get all products
func GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	priceMin, _ := strconv.ParseFloat(c.Query("priceMin"), 64)
	priceMax, _ := strconv.ParseFloat(c.Query("priceMax"), 64)
	stockMin, _ := strconv.Atoi(c.Query("stockMin"))
	stockMax, _ := strconv.Atoi(c.Query("stockMax"))

	params := repositories.ProductQueryParams{
		Search:    c.Query("search"),
		SortBy:    c.DefaultQuery("sortBy", "created_at"),
		SortOrder: c.DefaultQuery("sortOrder", "desc"),
		Page:      page,
		Limit:     limit,
	}

	// Optional filters
	if c.Query("priceMin") != "" {
		params.PriceMin = &priceMin
	}

	if c.Query("priceMax") != "" {
		params.PriceMax = &priceMax
	}

	if c.Query("stockMin") != "" {
		params.StockMin = &stockMin
	}

	if c.Query("stockMax") != "" {
		params.StockMax = &stockMax
	}

	products, total, err := services.GetAllProducts(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch products",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Product fetched successfully",
		"status":   true,
		"products": products,
		"pagination": gin.H{
			"total":      total,
			"page":       page,
			"limit":      limit,
			"totalPages": int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

// 3. Function to get product by id
func GetProduct(c *gin.Context) {
	idParams := c.Param("id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
			"status":  false,
		})
		return
	}

	product, err := services.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product fetched successfully",
		"status":  true,
		"product": product,
	})
}
