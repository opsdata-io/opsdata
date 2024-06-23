package utils

import (
	"strconv"
)

// ParseUint parses a string to uint.
func ParseUint(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 64)
	return uint(u)
}

// ParseInt parses a string to int.
func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
