package utils

import (
	"strconv"
)

// ParseUint parses a string to uint.
func ParseUint(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 64)
	return uint(u)
}
