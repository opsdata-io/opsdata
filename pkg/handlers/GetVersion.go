package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/version"

	_ "github.com/swaggo/fiber-swagger"
)

// GetVersion handles fetching version information
// @Summary Get version information
// @Description Retrieves a JSON response with version information
// @Tags Version
// @Produce json
// @Success 200 {object} map[string]interface{} "Successful operation"
// @Router /v1/version [get]
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
