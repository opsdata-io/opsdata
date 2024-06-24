package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// CreateUser handles creating a new user based on request payload
// @Summary Create a new user
// @Description Creates a new user based on request payload
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/users [post]
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid request payload"})
	}
	if err := utils.CreateUser(&user); err != nil { // Replace with your function to create a new user in database
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to create user"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
