package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/handlers"
	"github.com/opsdata-io/opsdata/backend/pkg/middleware"

	swagger "github.com/swaggo/fiber-swagger"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App) {
	// Customer routes
	app.Post("/v1/customers", middleware.AuthenticateJWTAndAPIKey, handlers.CreateCustomer)        // Create a new customer
	app.Get("/v1/customers", middleware.AuthenticateJWTAndAPIKey, handlers.GetCustomers)           // Get all customers
	app.Get("/v1/customers/search", middleware.AuthenticateJWTAndAPIKey, handlers.SearchCustomers) // Search customers
	app.Get("/v1/customers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.GetCustomer)        // Get a single customer by ID
	app.Put("/v1/customers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.UpdateCustomer)     // Update a customer by ID
	app.Delete("/v1/customers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.DeleteCustomer)  // Delete a customer by ID

	// User routes
	app.Get("/v1/users", middleware.AuthenticateJWTAndAPIKey, handlers.GetUsers)           // Get all users
	app.Get("/v1/users/:id", middleware.AuthenticateJWTAndAPIKey, handlers.GetUser)        // Get a single user by ID
	app.Post("/v1/users", middleware.AuthenticateJWTAndAPIKey, handlers.CreateUser)        // Create a new user
	app.Put("/v1/users/:id", middleware.AuthenticateJWTAndAPIKey, handlers.UpdateUser)     // Update a user by ID
	app.Delete("/v1/users/:id", middleware.AuthenticateJWTAndAPIKey, handlers.DeleteUser)  // Delete a user by ID
	app.Get("/v1/users/search", middleware.AuthenticateJWTAndAPIKey, handlers.SearchUsers) // Search users

	// Server routes
	app.Get("/v1/servers", middleware.AuthenticateJWTAndAPIKey, handlers.GetServers)           // Get all servers
	app.Post("/v1/servers", middleware.AuthenticateJWTAndAPIKey, handlers.CreateServer)        // Create a new server
	app.Get("/v1/servers", middleware.AuthenticateJWTAndAPIKey, handlers.GetServers)           // Get all servers
	app.Get("/v1/servers/search", middleware.AuthenticateJWTAndAPIKey, handlers.SearchServers) // Search servers
	app.Get("/v1/servers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.GetServer)        // Get a single server by ID
	app.Put("/v1/servers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.UpdateServer)     // Update a server by ID
	app.Delete("/v1/servers/:id", middleware.AuthenticateJWTAndAPIKey, handlers.DeleteServer)  // Delete a server by ID

	// Other endpoints
	app.Post("/v1/login", handlers.Login)
	app.Post("/v1/create-link", middleware.AuthenticateJWTAndAPIKey, handlers.CreateUploadLink)
	app.Post("/v1/upload/:link", middleware.AuthenticateJWTAndAPIKey, handlers.UploadFile)
	app.Get("/v1/files", middleware.AuthenticateJWTAndAPIKey, handlers.DownloadFiles)

	// Health check and version endpoints
	app.Get("/v1/version", handlers.GetVersion)
	app.Get("/v1/healthz", handlers.GetHealth)
	app.Get("/v1/readyz", handlers.GetReady)

	// Swagger documentation
	app.Get("/v1/swagger/*", swagger.FiberWrapHandler(swagger.URL("doc.json")))
}
