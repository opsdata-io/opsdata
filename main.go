package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/opsdata-io/opsdata/pkg/config"
	"github.com/opsdata-io/opsdata/pkg/metrics"
	"github.com/opsdata-io/opsdata/pkg/routes"
	"github.com/opsdata-io/opsdata/pkg/utils"

	// Import swag
	_ "github.com/opsdata-io/opsdata/docs"
	swagger "github.com/swaggo/fiber-swagger"
)

func main() {
	// Load configuration
	config.LoadConfiguration()

	fmt.Println("Starting OpsData API Server")

	// Initialize S3
	if err := utils.InitS3(); err != nil {
		fmt.Println("Failed to initialize S3:", err)
	}

	// Start metrics server
	go metrics.StartMetricsServer()

	// Setup Fiber
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path} - ${latency} - IP: ${reqHeader:X-Forwarded-For} - Host: ${host} - Protocol: ${protocol} - Referer: ${referer} - UserAgent: ${ua}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Local",
	}))

	// Routes
	routes.SetupRoutes(app)

	// Serve Swagger UI at the root path
	app.Get("/", swagger.FiberWrapHandler(swagger.URL("doc.json")))

	// Connect to Database
	utils.ConnectDB()

	// Start Server
	if err := app.Listen(fmt.Sprintf(":%d", config.CFG.ServerPort)); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
