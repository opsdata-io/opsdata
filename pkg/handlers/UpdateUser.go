package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// UpdateUser handles updating a user by ID
// @Summary Update a user by ID
// @Description Updates a user's information based on their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateUser models.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid request payload"})
	}
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "User ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	// Validate the id string to uint conversion
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid user ID"})
	}
	user, err := utils.GetUserByID(idUint) // Fetch the user from database by ID
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": "User not found"})
	}

	// Update the user fields
	user.Username = updateUser.Username
	user.Email = updateUser.Email

	if err := utils.UpdateUser(idUint, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to update user"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
