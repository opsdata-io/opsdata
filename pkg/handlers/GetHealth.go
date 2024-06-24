package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetHealth handles fetching the health status of the application
// @Summary Get Health Status
// @Description Retrieves the health status of the application
// @Tags Health
// @Produce plain
// @Success 200 {string} string "ok"
// @Router /v1/healthz [get]
func GetHealth(c *fiber.Ctx) error {
	return c.SendString("ok")
}
