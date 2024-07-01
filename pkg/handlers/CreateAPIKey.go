// handlers/CreateAPIKey.go

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// CreateAPIKey creates a new API key in the database and responds with the generated keys
// @Summary Create a new API key
// @Description Creates a new API key in the database
// @Tags API Keys
// @Accept json
// @Produce json
// @Param apikey body models.APIKey true "API Key object"
// @Success 201 {object} models.APIKey
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/apikeys [post]
func CreateAPIKey(c *fiber.Ctx) error {
	// Parse the request body into an APIKey struct
	var apiKey models.APIKey
	if err := c.BodyParser(&apiKey); err != nil {
		// Return bad request if request payload is invalid
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid request payload"})
	}

	// Validate required fields
	if apiKey.CustomerID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "CustomerID is required"})
	}

	if apiKey.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "UserID is required"})
	}

	// Create API key in the database
	if err := utils.CreateAPIKey(&apiKey); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to create API key"})
	}

	// Return the created API key with the generated keys
	return c.Status(fiber.StatusCreated).JSON(apiKey)
}
