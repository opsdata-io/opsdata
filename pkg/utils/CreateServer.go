package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// CreateServer creates a new server in the database
func CreateServer(server *models.Server) error {
	return DB.Create(server).Error
}
