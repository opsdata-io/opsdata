package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/models"
)

// VerifyAPIKey verifies the provided API accessKey and secretKey
func VerifyAPIKey(accessKey, secretKey string) (bool, *models.User, error) {
	// Lookup key by accessKey
	key, err := GetAPIKeyByAccessKey(accessKey)
	if err != nil {
		return false, nil, err
	}

	// Verify secret key
	err = VerifySecretKey(key.SecretKeyHash, secretKey)
	if err != nil {
		return false, nil, err
	}

	// Retrieve associated user
	user, err := GetUserByID(key.UserID)
	if err != nil {
		return false, nil, err
	}

	// API key and secretKey are valid
	return true, user, nil
}
