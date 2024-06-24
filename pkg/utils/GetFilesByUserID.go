package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// GetFilesByUserID returns all files from the database by user ID
func GetFilesByUserID(userID uint) ([]models.FileMetadata, error) {
	var files []models.FileMetadata
	result := DB.Where("user_id = ?", userID).Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}
	return files, nil
}
