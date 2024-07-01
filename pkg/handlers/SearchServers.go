package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// SearchServers handles searching servers based on a query parameter in the database
// @Summary Search servers by query parameter
// @Description Searches servers in the database based on a query parameter
// @Tags Servers
// @Param q query string true "Search query"
// @Produce json
// @Success 200 {array} models.Server
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/servers/search [get]
func SearchServers(c *fiber.Ctx) error {
	query := c.Query("q") // Get the 'q' query parameter for search query
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Search query 'q' is required"})
	}

	// Search servers in the database based on query
	servers, err := utils.SearchServers(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to search servers"})
	}

	// Return search results as JSON response
	return c.JSON(servers)
}
