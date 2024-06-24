package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/handlers"
	"github.com/opsdata-io/opsdata/pkg/middleware"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App) {
	// Customer routes
	app.Get("/v1/customers", handlers.GetCustomers)           // Get all customers
	app.Get("/v1/customers/:id", handlers.GetCustomer)        // Get a single customer by ID
	app.Post("/v1/customers", handlers.CreateCustomer)        // Create a new customer
	app.Put("/v1/customers/:id", handlers.UpdateCustomer)     // Update a customer by ID
	app.Delete("/v1/customers/:id", handlers.DeleteCustomer)  // Delete a customer by ID
	app.Get("/v1/customers/search", handlers.SearchCustomers) // Search customers

	// User routes
	app.Get("/v1/users", handlers.GetUsers)           // Get all users
	app.Get("/v1/users/:id", handlers.GetUser)        // Get a single user by ID
	app.Post("/v1/users", handlers.CreateUser)        // Create a new user
	app.Put("/v1/users/:id", handlers.UpdateUser)     // Update a user by ID
	app.Delete("/v1/users/:id", handlers.DeleteUser)  // Delete a user by ID
	app.Get("/v1/users/search", handlers.SearchUsers) // Search users

	// Other endpoints
	app.Post("/login", handlers.Login)
	app.Post("/create-link", middleware.AuthenticateJWT, handlers.CreateUploadLink)
	app.Post("/upload/:link", middleware.AuthenticateJWT, handlers.UploadFile)
	app.Get("/files", middleware.AuthenticateJWT, handlers.DownloadFiles)

	// Health check and version endpoints
	app.Get("/v1/version", handlers.GetVersion)
	app.Get("/v1/healthz", handlers.GetHealth)
	app.Get("/v1/readyz", handlers.GetReady)
}
