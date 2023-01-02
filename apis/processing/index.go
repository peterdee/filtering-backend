package processing

import "github.com/gofiber/fiber/v2"

func Initialize(app *fiber.App) {
	api := app.Group("/api/processing")

	api.Post("/", processImageController)
}
