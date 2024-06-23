package handlers

import (
	"github.com/gofiber/fiber/v2"

	_ "github.com/swaggo/fiber-swagger"
)

// @Summary Get Ready Status
// @Description Retrieves the readiness status of the application
// @Tags Ready
// @Produce json
// @Success 200 {object} map[string]interface{} "Successful operation"
// @Router /v1/ready [get]
func GetReady(c *fiber.Ctx) error {
	// Construct JSON response
	response := map[string]interface{}{"status": "ready"}

	// Return the readiness status as JSON response
	return c.JSON(response)
}
