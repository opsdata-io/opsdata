// utils/apikey.go

package utils

import (
	"crypto/rand"
)

// secureRandomSeed generates a secure random seed for rand package
func secureRandomSeed() int64 {
	var seed int64
	_, err := rand.Read([]byte{byte(seed)})
	if err != nil {
		panic(err)
	}
	return seed
}
