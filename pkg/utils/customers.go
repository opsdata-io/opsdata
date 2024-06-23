// utils/customers.go

package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	result := DB.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func GetCustomerByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	result := DB.First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func CreateCustomer(customer *models.Customer) error {
	return DB.Create(customer).Error
}

func UpdateCustomer(id uint, customer *models.Customer) error {
	rest := DB.First(&models.Customer{}, id)
	if rest.Error != nil {
		return rest.Error
	}
	return DB.Save(customer).Error
}

func DeleteCustomer(id uint) error {
	return DB.Delete(&models.Customer{}, id).Error
}

func SearchCustomers(query string) ([]models.Customer, error) {
	var customers []models.Customer
	if err := DB.Where("companyName LIKE ? OR contactName LIKE ?", "%"+query+"%", "%"+query+"%").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
