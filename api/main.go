package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mattmattox/opsdata/database"
	"github.com/mattmattox/opsdata/routes"
	"gorm.io/gorm/logger"
)

func main() {

	database.Connect()

	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
