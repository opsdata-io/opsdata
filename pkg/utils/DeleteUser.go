package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// DeleteUser deletes a user from the database
func DeleteUser(id uint) error {
	return DB.Delete(&models.User{}, id).Error
}
