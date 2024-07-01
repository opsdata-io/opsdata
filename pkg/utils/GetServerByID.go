package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// GetServerByID retrieves a server from the database by its ID
func GetServerByID(id uint) (*models.Server, error) {
	var server models.Server
	result := DB.First(&server, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &server, nil
}
