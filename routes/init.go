package routes

import (
	"estef.link/handlers"
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(app *fiber.App) {
	app.Get("/healthcheck", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// API
	api := app.Group("/api")
	api.Post("/shorten", handlers.ShortenURL)
}
