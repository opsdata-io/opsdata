package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// SearchCustomers handles searching customers based on a query parameter in the database
// @Summary Search customers by query parameter
// @Description Searches customers in the database based on a query parameter
// @Tags Customers
// @Param q query string true "Search query"
// @Produce json
// @Success 200 {array} models.Customer
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/customers/search [get]
func SearchCustomers(c *fiber.Ctx) error {
	query := c.Query("q") // Get the 'q' query parameter for search query
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Search query 'q' is required"})
	}

	// Search customers in the database based on query
	customers, err := utils.SearchCustomers(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to search customers"})
	}

	// Return search results as JSON response
	return c.JSON(customers)
}
