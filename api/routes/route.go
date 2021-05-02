package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattmattox/opsdata/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Healthz)
	app.Get("/healthz", controllers.Healthz)
	app.Get("/api/healthz", controllers.Healthz)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)
	app.Post("/api/key/create", controllers.CreateApiKey)
	app.Get("/api/key/verify", controllers.VerifyApiKey)
}
