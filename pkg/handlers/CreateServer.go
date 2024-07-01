package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// CreateServer creates a new server in the database and returns the new server object in JSON format
// @Summary Create a new server
// @Description Creates a new server in the database
// @Tags Servers
// @Accept json
// @Produce json
// @Param server body models.Server true "Server object"
// @Success 201 {object} models.Server
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/servers [post]
func CreateServer(c *fiber.Ctx) error {
	// Parse the request body into a Server struct
	var server models.Server
	if err := c.BodyParser(&server); err != nil {
		// Return bad request if request payload is invalid
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid request payload"})
	}

	// Validate required fields
	if server.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Name is required"})
	}
	if server.CustomerID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "CustomerID is required"})
	}
	// Optionally, validate other fields like IPAddress, Description, etc.

	// Set default values or perform any additional validation if needed
	server.LastPing = time.Now()

	// Create server in the database
	if err := utils.CreateServer(&server); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to create server"})
	}

	// Return the created server as JSON response
	return c.Status(fiber.StatusCreated).JSON(server)
}
