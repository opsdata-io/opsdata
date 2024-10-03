// utils/GetAPIKeyByAccessKey.go

package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// GetAPIKeyByAccessKey retrieves an API key from the database by its access key
func GetAPIKeyByAccessKey(accessKey string) (*models.APIKey, error) {
	var apiKey models.APIKey
	result := DB.Where("access_key = ?", accessKey).First(&apiKey)
	if result.Error != nil {
		return nil, result.Error
	}
	return &apiKey, nil
}
