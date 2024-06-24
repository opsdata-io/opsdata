package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// DeleteUser handles deleting a user by ID
// @Summary Delete a user by ID
// @Description Deletes a user based on their ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{} "User ID is required"
// @Failure 400 {object} map[string]interface{} "Invalid user ID"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "User ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	// Validate the id string to uint conversion
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid user ID"})
	}
	if err := utils.DeleteUser(idUint); err != nil { // Replace with your function to delete a user from database
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": "User not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
