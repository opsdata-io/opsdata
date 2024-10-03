package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/opsdata-io/opsdata/backend/pkg/config"
	"github.com/opsdata-io/opsdata/backend/pkg/metrics"
	"github.com/opsdata-io/opsdata/backend/pkg/routes"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
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
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	// Custom logging middleware
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path} - ${latency} - IP: ${reqHeader:CF-Connecting-IP} - Host: ${host} - Protocol: ${protocol} - Referer: ${referer} - UserAgent: ${ua}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Local",
	}))

	// Routes
	routes.SetupRoutes(app)

	// Connect to Database
	utils.ConnectDB()

	// Start Server
	if err := app.Listen(fmt.Sprintf(":%d", config.CFG.ServerPort)); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
