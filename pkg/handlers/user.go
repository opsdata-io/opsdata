package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models" // Assuming you have a models package with user model
	"github.com/opsdata-io/opsdata/pkg/utils"  // Assuming you have utility functions for database operations
)

// GetUsers handles fetching all users
func GetUsers(c *fiber.Ctx) error {
	users, err := utils.GetAllUsers() // Replace with your function to fetch all users from database
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

// GetUser handles fetching a single user by ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user, err := utils.GetUserByID(idUint) // Replace with your function to fetch a user by ID from database
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// CreateUser handles creating a new user
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	if err := utils.CreateUser(&user); err != nil { // Replace with your function to create a new user in database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// UpdateUser handles updating a user by ID
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateUser models.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	// Validate the id string to uint conversion
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	user, err := utils.GetUserByID(idUint) // Fetch the user from database by ID
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Update the user fields
	user.Username = updateUser.Username
	user.Email = updateUser.Email

	if err := utils.UpdateUser(idUint, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// DeleteUser handles deleting a user by ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is required"})
	}
	// Parse the id string to uint
	idUint := utils.ParseUint(id)
	// Validate the id string to uint conversion
	if idUint == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	if err := utils.DeleteUser(idUint); err != nil { // Replace with your function to delete a user from database
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// SearchUsers handles searching users based on certain criteria
func SearchUsers(c *fiber.Ctx) error {
	// Implement search logic based on query parameters or request body
	// Example: search by name, email, etc.
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": "Search functionality not implemented yet"})
}
