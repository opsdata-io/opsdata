package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// UpdateServer updates a server in the database
func UpdateServer(id uint, server *models.Server) error {
	rest := DB.First(&models.Server{}, id)
	if rest.Error != nil {
		return rest.Error
	}
	return DB.Save(server).Error
}
