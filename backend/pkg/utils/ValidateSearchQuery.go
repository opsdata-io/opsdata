package utils

import (
	"errors"
	"regexp"
)

// ValidateSearchQuery validates the search query based on DNS rules
func ValidateSearchQuery(query string) error {
	// Define the regular expression for validating the query
	// Allows alphanumeric characters and hyphens, but not at the beginning or end
	var validQueryRegex = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?$`)

	// Check if the query matches the regular expression
	if !validQueryRegex.MatchString(query) {
		return errors.New("invalid query: must be alphanumeric with optional hyphens, not at the beginning or end")
	}
	return nil
}
