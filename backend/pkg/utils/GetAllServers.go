package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// GetAllServers returns all servers from the database
func GetAllServers() ([]models.Server, error) {
	var servers []models.Server
	result := DB.Find(&servers)
	if result.Error != nil {
		return nil, result.Error
	}
	return servers, nil
}
