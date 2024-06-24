package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// CreateCustomer creates a new customer in the database
func CreateCustomer(customer *models.Customer) error {
	return DB.Create(customer).Error
}
