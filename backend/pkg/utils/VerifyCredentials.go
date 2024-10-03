package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// VerifyCredentials checks if the email and password are valid and returns true if they are valid, otherwise it returns false
func VerifyCredentials(email, password string) bool {
	var user models.User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
