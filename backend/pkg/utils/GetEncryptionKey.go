package utils

import (
	"github.com/opsdata-io/opsdata/backend/pkg/config"
)

// GetEncryptionKey returns the encryption key
func GetEncryptionKey() string {
	return config.CFG.EncryptionKey
}
