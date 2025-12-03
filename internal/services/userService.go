package services

import (
	"errors"
	"strconv"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

// 1. Function to register user
func RegisterUser(user *models.User) error {
	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save user
	return repositories.CreateUser(user)
}

// 2. Function to update user profile
type UpdateProfileInput struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Email     *string `json:"email,omitempty"`
}

func UpdateUserProfile(user *models.User, input *UpdateProfileInput) (*models.User, error) {
	// Only update fields if they are provided
	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		user.LastName = *input.LastName
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	// Save via repository
	if err := repositories.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// 3. Function to change password
type ChangePasswordInput struct {
	OldPassword string `json:"oldPassword" binding:"required,min=6"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}

func ChangePassword(user *models.User, input *ChangePasswordInput) error {
	// Verify Old Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		return errors.New("old password is incorrect")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update user password
	user.Password = string(hashedPassword)
	return repositories.UpdateUser(user)
}

// 4. Function to get all users
func GetAllUsers(pageStr, limitStr, search, role, sortBy, sortOrder string) ([]models.User, int64, error) {
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	return repositories.GetAllUsers(offset, limit, search, role, sortBy, sortOrder)
}

// 5. Function to delete user by id
func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}
