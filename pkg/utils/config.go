package utils

import (
	"github.com/opsdata-io/opsdata/pkg/config"
)

func GetJWTSecret() string {
	return config.CFG.JWTSecret
}

func GetEncryptionKey() string {
	return config.CFG.EncryptionKey
}
