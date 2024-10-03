package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// DeleteServer handles deleting a server by ID from the database
// @Summary Delete a server by ID
// @Description Deletes a server from the database by ID
// @Tags Servers
// @Param id path string true "Server ID"
// @Produce json
// @Success 200 {object} map[string]interface{} "Server successfully deleted"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/servers/{id} [delete]
func DeleteServer(c *fiber.Ctx) error {
	id := c.Params("id")

	// Parse ID into uint
	serverID := utils.ParseUint(id)

	// Delete server from the database
	if err := utils.DeleteServer(serverID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to delete server"})
	}

	// Return success message if server deletion is successful
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"data": "Server successfully deleted"})
}
