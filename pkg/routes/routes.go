package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/handlers"
	"github.com/opsdata-io/opsdata/pkg/middleware"

	swagger "github.com/swaggo/fiber-swagger"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App) {
	// Customer routes
	app.Post("/api/customers", middleware.AuthenticateJWTAndAPIKey, handlers.CreateCustomer)        // Create a new customer
	app.Get("/api/customers", middleware.AuthenticateJWTAndAPIKey, handlers.GetCustomers)           // Get all customers
	app.Get("/api/customers/search", middleware.AuthenticateJWTAndAPIKey, handlers.SearchCustomers) // Search customers
	app.Get("/api/customers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.GetCustomer)        // Get a single customer by ID
	app.Put("/api/customers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.UpdateCustomer)     // Update a customer by ID
	app.Delete("/api/customers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.DeleteCustomer)  // Delete a customer by ID

	// User routes
	app.Get("/api/users", middleware.AuthenticateJWTAndAPIKey, handlers.GetUsers)           // Get all users
	app.Get("/api/users/:id", middleware.AuthenticateJWTAndAPIKey, handlers.GetUser)        // Get a single user by ID
	app.Post("/api/users", middleware.AuthenticateJWTAndAPIKey, handlers.CreateUser)        // Create a new user
	app.Put("/api/users/:id", middleware.AuthenticateJWTAndAPIKey, handlers.UpdateUser)     // Update a user by ID
	app.Delete("/api/users/:id", middleware.AuthenticateJWTAndAPIKey, handlers.DeleteUser)  // Delete a user by ID
	app.Get("/api/users/search", middleware.AuthenticateJWTAndAPIKey, handlers.SearchUsers) // Search users

	// Server routes
	app.Get("/api/servers", middleware.AuthenticateJWTAndAPIKey, handlers.GetServers)           // Get all servers
	app.Post("/api/servers", middleware.AuthenticateJWTAndAPIKey, handlers.CreateServer)        // Create a new server
	app.Get("/api/servers", middleware.AuthenticateJWTAndAPIKey, handlers.GetServers)           // Get all servers
	app.Get("/api/servers/search", middleware.AuthenticateJWTAndAPIKey, handlers.SearchServers) // Search servers
	app.Get("/api/servers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.GetServer)        // Get a single server by ID
	app.Put("/api/servers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.UpdateServer)     // Update a server by ID
	app.Delete("/api/servers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.DeleteServer)  // Delete a server by ID

	// Other endpoints
	app.Post("/api/login", handlers.Login)
	app.Post("/api/create-link", middleware.AuthenticateJWTAndAPIKey, handlers.CreateUploadLink)
	app.Post("/api/upload/:link", middleware.AuthenticateJWTAndAPIKey, handlers.UploadFile)
	app.Get("/api/files", middleware.AuthenticateJWTAndAPIKey, handlers.DownloadFiles)

	// Health check and version endpoints
	app.Get("/api/version", handlers.GetVersion)
	app.Get("/api/healthz", handlers.GetHealth)
	app.Get("/api/readyz", handlers.GetReady)

	// Swagger documentation
	app.Get("/api/swagger/*", swagger.FiberWrapHandler(swagger.URL("doc.json")))
}
