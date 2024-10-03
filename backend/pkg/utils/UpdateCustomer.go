package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// UpdateCustomer updates a customer in the database
func UpdateCustomer(id uint, customer *models.Customer) error {
	rest := DB.First(&models.Customer{}, id)
	if rest.Error != nil {
		return rest.Error
	}
	return DB.Save(customer).Error
}
