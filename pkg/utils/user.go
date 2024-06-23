package utils

import (
	"github.com/opsdata-io/opsdata/pkg/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

func UpdateUser(id uint, user *models.User) error {
	result := DB.First(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return DB.Delete(&models.User{}, id).Error
}

func SearchUsers(query string) ([]models.User, error) {
	var users []models.User
	if err := DB.Where("username LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
