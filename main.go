package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/opsdata-io/opsdata/pkg/config"
	"github.com/opsdata-io/opsdata/pkg/metrics"
	"github.com/opsdata-io/opsdata/pkg/routes"
	"github.com/opsdata-io/opsdata/pkg/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	config.LoadConfiguration()

	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	logger.Infoln("Starting OpsData API Server")

	// Initialize S3
	if err := utils.InitS3(); err != nil {
		logger.Fatalln("Failed to initialize S3:", err)
	}

	// Start metrics server
	go metrics.StartMetricsServer()

	// Setup Fiber
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	// Custom logging middleware
	app.Use(func(c *fiber.Ctx) error {
		logger := logrus.New()
		logger.SetFormatter(&logrus.TextFormatter{})
		logger.Infof("%s %s", c.Method(), c.Path())
		return c.Next()
	})

	// Routes
	routes.SetupRoutes(app)

	// Connect to Database
	utils.ConnectDB()

	// Start Server
	if err := app.Listen(fmt.Sprintf(":%d", config.CFG.ServerPort)); err != nil {
		logger.Fatalln("Failed to start server:", err)
	}
}
