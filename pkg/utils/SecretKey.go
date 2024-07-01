package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashSecretKey hashes the secret key using bcrypt and returns the hashed password.
func HashSecretKey(secretKey string) (string, error) {
	hashedSecretKey, err := bcrypt.GenerateFromPassword([]byte(secretKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedSecretKey), nil
}

// VerifySecretKey verifies if the provided secret key matches the hashed password.
func VerifySecretKey(hashedSecretKey, secretKey string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSecretKey), []byte(secretKey))
}
