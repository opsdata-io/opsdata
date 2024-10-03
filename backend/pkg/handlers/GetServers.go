package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// GetServers handles fetching all servers from the database.
// @Summary Get all servers
// @Description Retrieves all servers from the database
// @Tags Servers
// @Produce json
// @Success 200 {array} models.Server
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/servers [get]
func GetServers(c *fiber.Ctx) error {
	// Retrieve all servers
	servers, err := utils.GetAllServers()
	if err != nil {
		// Return internal server error if fetching servers fails
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to fetch servers"})
	}
	// Return servers as JSON response
	return c.JSON(servers)
}
