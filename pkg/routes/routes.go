package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/handlers"
	"github.com/opsdata-io/opsdata/pkg/middleware"

	_ "github.com/opsdata-io/opsdata/docs"
	swagger "github.com/swaggo/fiber-swagger"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App) {
	// Customer routes
	app.Post("/api/customers", middleware.AuthenticateJWT, handlers.CreateCustomer)        // Create a new customer
	app.Get("/api/customers", middleware.AuthenticateJWT, handlers.GetCustomers)           // Get all customers
	app.Get("/api/customers/search", middleware.AuthenticateJWT, handlers.SearchCustomers) // Search customers
	app.Get("/api/customers/:id", middleware.AuthenticateJWT, handlers.GetCustomer)        // Get a single customer by ID
	app.Put("/api/customers/:id", middleware.AuthenticateJWT, handlers.UpdateCustomer)     // Update a customer by ID
	app.Delete("/api/customers/:id", middleware.AuthenticateJWT, handlers.DeleteCustomer)  // Delete a customer by ID

	// User routes
	app.Get("/api/users", middleware.AuthenticateJWT, handlers.GetUsers)           // Get all users
	app.Get("/api/users/:id", middleware.AuthenticateJWT, handlers.GetUser)        // Get a single user by ID
	app.Post("/api/users", middleware.AuthenticateJWT, handlers.CreateUser)        // Create a new user
	app.Put("/api/users/:id", middleware.AuthenticateJWT, handlers.UpdateUser)     // Update a user by ID
	app.Delete("/api/users/:id", middleware.AuthenticateJWT, handlers.DeleteUser)  // Delete a user by ID
	app.Get("/api/users/search", middleware.AuthenticateJWT, handlers.SearchUsers) // Search users

	// Other endpoints
	app.Post("/api/login", handlers.Login)
	app.Post("/api/create-link", middleware.AuthenticateJWT, handlers.CreateUploadLink)
	app.Post("/api/upload/:link", middleware.AuthenticateJWT, handlers.UploadFile)
	app.Get("/files", middleware.AuthenticateJWT, handlers.DownloadFiles)

	// Health check and version endpoints
	app.Get("/api/version", handlers.GetVersion)
	app.Get("/api/healthz", handlers.GetHealth)
	app.Get("/api/readyz", handlers.GetReady)

	// Swagger documentation
	app.Get("/swagger/*", swagger.FiberWrapHandler(swagger.URL("doc.json")))
}
