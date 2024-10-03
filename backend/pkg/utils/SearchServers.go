package utils

import (
	"strings"

	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// SearchServers searches for servers in the database without case sensitivity
func SearchServers(query string) ([]models.Server, error) {
	// Validate the search query
	if err := ValidateSearchQuery(query); err != nil {
		return nil, err
	}

	// Perform case-insensitive search
	var servers []models.Server
	query = strings.ToLower(query)
	if err := DB.Where("LOWER(name) LIKE ? OR LOWER(device_type) LIKE ? OR LOWER(ip_address) LIKE ? OR LOWER(description) LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&servers).Error; err != nil {
		return nil, err
	}
	return servers, nil
}
