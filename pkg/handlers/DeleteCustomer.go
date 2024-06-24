package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// DeleteCustomer handles deleting a customer by ID from the database
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
