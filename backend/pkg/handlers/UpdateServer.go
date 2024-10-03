package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/models"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// UpdateServer handles updating an existing server by ID
// @Summary Update a server by ID
// @Description Updates an existing server in the database by ID
// @Tags Servers
// @Param id path string true "Server ID"
// @Accept json
// @Produce json
// @Param server body models.Server true "Updated server object"
// @Success 200 {object} models.Server
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Server not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/servers/{id} [put]
func UpdateServer(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateServer models.Server
	if err := c.BodyParser(&updateServer); err != nil {
		// Return bad request if request payload is invalid
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid request payload"})
	}
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Server ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	// Validate the id string to uint conversion
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid server ID"})
	}
	// Retrieve existing server by ID
	server, err := utils.GetServerByID(idUint)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": "Server not found"})
	}

	// Update server fields
	server.Name = updateServer.Name
	server.DeviceType = updateServer.DeviceType
	server.IPAddress = updateServer.IPAddress
	server.Description = updateServer.Description
	server.LastPing = updateServer.LastPing

	// Update server in the database
	if err := utils.UpdateServer(idUint, server); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to update server"})
	}
	// Return updated server as JSON response
	return c.Status(fiber.StatusOK).JSON(server)
}
