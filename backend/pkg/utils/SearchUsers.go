package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// SearchUsers searches for users in the database
func SearchUsers(query string) ([]models.User, error) {
	var users []models.User
	if err := DB.Where("username LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
