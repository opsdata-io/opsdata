package utils

import (
	"strings"

	"github.com/opsdata-io/opsdata/pkg/models"
)

// SearchCustomers searches for customers in the database without case sensitivity
func SearchCustomers(query string) ([]models.Customer, error) {
	// Validate the search query
	if err := ValidateSearchQuery(query); err != nil {
		return nil, err
	}

	// Perform case-insensitive search
	var customers []models.Customer
	query = strings.ToLower(query)
	if err := DB.Where("LOWER(company_name) LIKE ?", "%"+query+"%").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
