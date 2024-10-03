package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// GetUsers handles fetching all users
// @Summary Get all users
// @Description Retrieves a list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/users [get]
func GetUsers(c *fiber.Ctx) error {
	users, err := utils.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to fetch users"})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
