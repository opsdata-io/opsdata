// utils/CreateAPIKey.go

package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessKeyLength = 8
	secretKeyLength = 24
)

// CreateAPIKey creates a new API key in the database
func CreateAPIKey(apiKey *models.APIKey) error {
	// Generate AccessKey
	accessKey, err := generateRandomString(accessKeyLength)
	if err != nil {
		return err
	}
	apiKey.AccessKey = accessKey

	// Generate SecretKey
	secretKey, err := generateRandomString(secretKeyLength)
	if err != nil {
		return err
	}

	// Hash the SecretKey using bcrypt
	hashedSecretKey, err := bcrypt.GenerateFromPassword([]byte(secretKey), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	apiKey.SecretKeyHash = string(hashedSecretKey)

	// Save the API key to the database
	return DB.Create(apiKey).Error
}
