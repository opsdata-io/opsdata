package utils

import (
	"github.com/opsdata-io/opsdata/pkg/config"
)

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() string {
	return config.CFG.JWTSecret
}
