package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// CreateUser creates a new user in the database
func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}