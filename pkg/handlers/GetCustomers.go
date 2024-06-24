package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// GetCustomers handles fetching all customers from the database.
// @Summary Get all customers
// @Description Retrieves all customers from the database
// @Tags Customers
// @Produce json
// @Success 200 {array} models.Customer
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/customers [get]
func GetCustomers(c *fiber.Ctx) error {
	// Retrieve all customers
	customers, err := utils.GetAllCustomers()
	if err != nil {
		// Return internal server error if fetching customers fails
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to fetch customers"})
	}
	// Return customers as JSON response
	return c.JSON(customers)
}
