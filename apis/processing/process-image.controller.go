package processing

import (
	"io"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/brille"

	"filtering/configuration"
)

func processImageController(context *fiber.Ctx) error {
	file, fileError := context.FormFile("image")
	filter := context.FormValue("filter")
	if fileError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	// get file extension
	partials := strings.Split(file.Filename, ".")
	fileExtension := partials[len(partials)-1]
	if fileExtension == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidData,
		)
	}

	// convert file into io.Reader
	fileHandle, readerError := file.Open()
	if readerError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	format := fileExtension
	var processingError error
	var result io.Reader
	if filter == "binary" {
		result, format, processingError = brille.Binary(
			fileHandle,
			127,
		)
	}
	if filter == "grayscale" {
		result, format, processingError = brille.Grayscale(
			fileHandle,
			brille.GRAYSCALE_AVERAGE,
		)
	}
	if processingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	context.Set("Content-Type", "image/"+format)
	// TODO: set filename if necessary
	// context.Set("Content-Disposition", "attachment; filename="filename.jpg"")
	return context.SendStream(result)
}
