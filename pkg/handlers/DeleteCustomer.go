package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"

	_ "github.com/swaggo/fiber-swagger"
)

// @Summary Delete a customer by ID
// @Description Deletes a customer from the database by ID
// @Tags Customers
// @Param id path string true "Customer ID"
// @Produce json
// @Success 200 {object} map[string]interface{} "Customer successfully deleted"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/customers/{id} [delete]
func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer
	customer.ID = utils.ParseUint(id)
	// Delete customer from the database
	if err := utils.DeleteCustomer(customer.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to delete customer"})
	}
	// Return success message if customer deletion is successful
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"data": "Customer successfully deleted"})
}

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
