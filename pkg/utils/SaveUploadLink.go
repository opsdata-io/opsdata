package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// SaveUploadLink saves the upload link to the database
func SaveUploadLink(link *models.UploadLink) error {
	return DB.Create(link).Error
}
