package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"
)

// ValidSubscriptionStatus defines the list of valid subscription status values
var ValidSubscriptionStatus = []string{"Active", "Inactive", "Pending"}

// isValidSubscriptionStatus checks if the provided subscription status is valid
func isValidSubscriptionStatus(status string) bool {
	for _, validStatus := range ValidSubscriptionStatus {
		if status == validStatus {
			return true
		}
	}
	return false
}

func GetCustomers(c *fiber.Ctx) error {
	customers, err := utils.GetAllCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch customers"})
	}
	return c.JSON(customers)
}

func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	customer, err := utils.GetCustomerByID(idUint)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}
	return c.JSON(customer)
}

// CreateCustomer handles the creation of a new customer.
func CreateCustomer(c *fiber.Ctx) error {
	fmt.Println("CreateCustomer")

	// Parse the request body into a Customer struct
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validate required fields
	if customer.CompanyName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "CompanyName is required"})
	}
	if customer.Address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Address is required"})
	}
	if customer.ContactName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ContactName is required"})
	}
	if customer.ContactTitle == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ContactTitle is required"})
	}
	if customer.ContactEmail == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ContactEmail is required"})
	}
	if customer.ContactPhone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ContactPhone is required"})
	}
	// Validate SubscriptionStatus
	if !isValidSubscriptionStatus(customer.SubscriptionStatus) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid SubscriptionStatus"})
	}

	if err := utils.CreateCustomer(&customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create customer"})
	}

	// Return the created customer as JSON response
	return c.Status(fiber.StatusCreated).JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateCustomer models.Customer
	if err := c.BodyParser(&updateCustomer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Customer ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	// Validate the id string to uint conversion
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer ID"})
	}
	customer, err := utils.GetCustomerByID(idUint)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}

	// Update the customer fields
	customer.CompanyName = updateCustomer.CompanyName
	customer.Address = updateCustomer.Address
	customer.ContactName = updateCustomer.ContactName
	customer.ContactTitle = updateCustomer.ContactTitle
	customer.ContactEmail = updateCustomer.ContactEmail
	customer.ContactPhone = updateCustomer.ContactPhone
	customer.Notes = updateCustomer.Notes
	customer.SubscriptionStatus = updateCustomer.SubscriptionStatus

	// Update the customer in the database
	if err := utils.UpdateCustomer(idUint, customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update customer"})
	}
	return c.Status(fiber.StatusOK).JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer
	customer.ID = utils.ParseUint(id)
	if err := utils.DeleteCustomer(customer.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete customer"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Customer successfully deleted"})
}

func SearchCustomers(c *fiber.Ctx) error {
	query := c.Query("q") // Get the 'q' query parameter for search query
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Search query 'q' is required"})
	}

	customers, err := utils.SearchCustomers(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to search customers"})
	}

	return c.JSON(customers)
}
