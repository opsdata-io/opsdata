package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// GetUserByID gets a user by ID from the database
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
