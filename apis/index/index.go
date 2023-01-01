package index

import "github.com/gofiber/fiber/v2"

func Initialize(app *fiber.App) {
	api := app.Group("/")

	api.Get("/", indexController)
	api.Get("/api", indexController)
}
