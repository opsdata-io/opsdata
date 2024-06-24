package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// DeleteCustomer deletes a customer from the database
func DeleteCustomer(id uint) error {
	return DB.Delete(&models.Customer{}, id).Error
}
