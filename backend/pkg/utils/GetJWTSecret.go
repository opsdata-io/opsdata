package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/config"
)

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() string {
	return config.CFG.JWTSecret
}
