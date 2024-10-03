package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/models"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// CreateCustomer creates a new customer in the database and returns the new customer object in JSON format
// @Summary Create a new customer
// @Description Creates a new customer in the database
// @Tags Customers
// @Accept json
// @Produce json
// @Param customer body models.Customer true "Customer object"
// @Success 201 {object} models.Customer
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/customers [post]
func CreateCustomer(c *fiber.Ctx) error {
	// Parse the request body into a Customer struct
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		// Return bad request if request payload is invalid
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid request payload"})
	}

	// Validate required fields
	if customer.CompanyName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "CompanyName is required"})
	}
	if customer.Address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Address is required"})
	}
	// Validate SubscriptionStatus
	if !isValidSubscriptionStatus(customer.SubscriptionStatus) {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid SubscriptionStatus"})
	}

	// Create customer in the database
	if err := utils.CreateCustomer(&customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to create customer"})
	}

	// Return the created customer as JSON response
	return c.Status(fiber.StatusCreated).JSON(customer)
}
