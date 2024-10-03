package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetReady handles fetching the ready status of the application
// @Summary Get Ready Status
// @Description Retrieves the readiness status of the application
// @Tags Ready
// @Produce plain
// @Success 200 {string} string "ok"
// @Router /v1/readyz [get]
func GetReady(c *fiber.Ctx) error {
	return c.SendString("ok")
}
