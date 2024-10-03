package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// GetCustomerByID retrieves a customer from the database by its ID
func GetCustomerByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	result := DB.First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}
