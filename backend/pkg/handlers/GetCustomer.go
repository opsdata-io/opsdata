package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// GetCustomer handles fetching a customer from the database by ID.
// @Summary Get a customer by ID
// @Description Retrieves a customer from the database by ID
// @Tags Customers
// @Param id path string true "Customer ID"
// @Produce json
// @Success 200 {object} models.Customer
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Customer not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/customers/{id} [get]
func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Customer ID is required"})
	}

	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid customer ID"})
	}

	// Retrieve customer by ID
	customer, err := utils.GetCustomerByID(idUint)
	if err != nil {
		// Return not found if customer with ID is not found
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": "Customer not found"})
	}

	// Return customer as JSON response
	return c.JSON(customer)
}
