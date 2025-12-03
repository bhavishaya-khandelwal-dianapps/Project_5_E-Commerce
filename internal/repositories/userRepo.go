package repositories

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
)

// 1. Function to create user
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// 2. Function to get user by id
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

// 3. Function to get user by email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// 4. Function to update user
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

// 5. Function to get get all users
func GetAllUsers(offset, limit int, search, role, sortBy, sortOrder string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := config.DB.Model(&models.User{})

	// Search by firstName, lastName and email
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?", searchPattern, searchPattern, searchPattern)
	}

	// Filter by role
	if role != "" {
		query = query.Where("role = ?", role)
	}

	// Count total matching users
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination + Sorting
	if err := query.Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
