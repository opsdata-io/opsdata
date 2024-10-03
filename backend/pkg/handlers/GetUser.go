package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// GetUser handles fetching a single user by ID
// @Summary Get a user by ID
// @Description Retrieves a user by their ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "User ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid user ID"})
	}
	user, err := utils.GetUserByID(idUint) // Replace with your function to fetch a user by ID from database
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
