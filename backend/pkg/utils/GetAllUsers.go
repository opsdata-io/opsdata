package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// GetAllUsers returns all users from the database
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
