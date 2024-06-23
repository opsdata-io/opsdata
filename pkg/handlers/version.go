package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/version"
)

// GetVersion handles the /api/version endpoint
func GetVersion(c *fiber.Ctx) error {
	// Fetch version information from the version package
	versionInfo := version.Version
	commit := version.GitCommit
	buildTime := version.BuildTime

	// Construct JSON response
	response := map[string]interface{}{
		"version":   versionInfo,
		"gitCommit": commit,
		"buildTime": buildTime,
	}

	// Return the version information as JSON response
	return c.JSON(response)
}
