package processing

import (
	"github.com/gofiber/fiber/v2"

	"filtering/utilities"
)

func processImageController(context *fiber.Ctx) error {
	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
