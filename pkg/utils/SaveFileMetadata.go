package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// SaveFileMetadata saves file metadata to the database
func SaveFileMetadata(file *models.FileMetadata) error {
	return DB.Create(file).Error
}
