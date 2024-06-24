package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// UpdateCustomer handles updating an existing customer by ID
// @Summary Update a customer by ID
// @Description Updates an existing customer in the database by ID
// @Tags Customers
// @Param id path string true "Customer ID"
// @Accept json
// @Produce json
// @Param customer body models.Customer true "Updated customer object"
// @Success 200 {object} models.Customer
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Customer not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/customers/{id} [put]
func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateCustomer models.Customer
	if err := c.BodyParser(&updateCustomer); err != nil {
		// Return bad request if request payload is invalid
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid request payload"})
	}
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Customer ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	// Validate the id string to uint conversion
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid customer ID"})
	}
	// Retrieve existing customer by ID
	customer, err := utils.GetCustomerByID(idUint)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": "Customer not found"})
	}

	// Update customer fields
	customer.CompanyName = updateCustomer.CompanyName
	customer.Address = updateCustomer.Address
	customer.ContactName = updateCustomer.ContactName
	customer.ContactTitle = updateCustomer.ContactTitle
	customer.ContactEmail = updateCustomer.ContactEmail
	customer.ContactPhone = updateCustomer.ContactPhone
	customer.Notes = updateCustomer.Notes
	customer.SubscriptionStatus = updateCustomer.SubscriptionStatus

	// Update customer in the database
	if err := utils.UpdateCustomer(idUint, customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to update customer"})
	}
	// Return updated customer as JSON response
	return c.Status(fiber.StatusOK).JSON(customer)
}
