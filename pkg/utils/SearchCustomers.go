package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

// SearchCustomers searches for customers in the database
func SearchCustomers(query string) ([]models.Customer, error) {
	var customers []models.Customer
	if err := DB.Where("companyName LIKE ? OR contactName LIKE ?", "%"+query+"%", "%"+query+"%").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
