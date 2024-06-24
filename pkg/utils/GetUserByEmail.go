package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// GetUserByEmail retrieves a user from the database by email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
