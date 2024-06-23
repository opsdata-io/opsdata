package handlers

import (
	"github.com/gofiber/fiber/v2"

	_ "github.com/swaggo/fiber-swagger"
)

// @Summary Get Health Status
// @Description Retrieves the health status of the application
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]interface{} "Successful operation"
// @Router /v1/health [get]
func GetHealth(c *fiber.Ctx) error {
	// Construct JSON response
	response := map[string]interface{}{"status": "ok"}

	// Return the health status as JSON response
	return c.JSON(response)
}
