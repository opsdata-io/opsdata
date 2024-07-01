package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// GetServer handles fetching a server from the database by ID.
// @Summary Get a server by ID
// @Description Retrieves a server from the database by ID
// @Tags Servers
// @Param id path string true "Server ID"
// @Produce json
// @Success 200 {object} models.Server
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Server not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/servers/{id} [get]
func GetServer(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Server ID is required"})
	}

	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid server ID"})
	}

	// Retrieve server by ID
	server, err := utils.GetServerByID(idUint)
	if err != nil {
		// Return not found if server with ID is not found
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": "Server not found"})
	}

	// Return server as JSON response
	return c.JSON(server)
}
