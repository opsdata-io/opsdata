package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// UpdateUser updates a user in the database
func UpdateUser(id uint, user *models.User) error {
	result := DB.First(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return DB.Save(user).Error
}
