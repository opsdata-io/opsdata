package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/handlers"
	"github.com/opsdata-io/opsdata/pkg/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Customer routes
	app.Get("/api/customers", handlers.GetCustomers)           // Get all customers
	app.Get("/api/customers/:id", handlers.GetCustomer)        // Get a single customer by ID
	app.Post("/api/customers", handlers.CreateCustomer)        // Create a new customer
	app.Put("/api/customers/:id", handlers.UpdateCustomer)     // Update a customer by ID
	app.Delete("/api/customers/:id", handlers.DeleteCustomer)  // Delete a customer by ID
	app.Get("/api/customers/search", handlers.SearchCustomers) // Search customers

	// User routes
	app.Get("/api/users", handlers.GetUsers)           // Get all users
	app.Get("/api/users/:id", handlers.GetUser)        // Get a single user by ID
	app.Post("/api/users", handlers.CreateUser)        // Create a new user
	app.Put("/api/users/:id", handlers.UpdateUser)     // Update a user by ID
	app.Delete("/api/users/:id", handlers.DeleteUser)  // Delete a user by ID
	app.Get("/api/users/search", handlers.SearchUsers) // Search users

	// Other endpoints
	app.Post("/login", handlers.Login)
	app.Post("/create-link", middleware.AuthenticateJWT, handlers.CreateUploadLink)
	app.Post("/upload/:link", middleware.AuthenticateJWT, handlers.UploadFile)
	app.Get("/files", middleware.AuthenticateJWT, handlers.DownloadFiles)

	// Version endpoint
	app.Get("/api/version", handlers.GetVersion)
}
