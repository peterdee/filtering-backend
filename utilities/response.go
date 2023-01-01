package utilities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/gohelpers"

	"filtering/configuration"
)

func Response(payload ResponsePayloadStruct) error {
	info := payload.Info
	if info == "" {
		info = configuration.RESPONSE_MESSAGES.OK
	}

	status := payload.Status
	if status == 0 {
		status = fiber.StatusOK
	}

	responseStruct := fiber.Map{
		"datetime": gohelpers.MakeTimestamp(),
		"info":     info,
		"request":  payload.Context.OriginalURL() + " [" + payload.Context.Method() + "]",
		"status":   status,
	}

	if payload.Data != nil {
		responseStruct["data"] = payload.Data
	}

	return payload.Context.Status(status).JSON(responseStruct)
}
