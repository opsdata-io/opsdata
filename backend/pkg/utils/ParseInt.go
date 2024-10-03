package utils

import (
	"strconv"
)

// ParseInt parses a string to int.
func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
