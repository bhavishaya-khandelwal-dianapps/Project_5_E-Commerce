package services

import (
	"errors"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(email string, password string) (string, *models.User, error) {
	var user models.User

	// Check email exist
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Generate Token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	return token, &user, nil
}
