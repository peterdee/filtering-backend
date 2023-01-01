package index

import (
	"github.com/gofiber/fiber/v2"

	"filtering/utilities"
)

func indexController(context *fiber.Ctx) error {
	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
