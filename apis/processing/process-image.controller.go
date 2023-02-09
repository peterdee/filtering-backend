package processing

import (
	"io"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/brille/v2"
	"github.com/julyskies/gohelpers"

	"filtering/configuration"
	"filtering/utilities"
)

func processImageController(context *fiber.Ctx) error {
	file, fileError := context.FormFile("image")
	filter := context.FormValue("filter")
	if fileError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if filter == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingFilterName,
		)
	}
	if !gohelpers.IncludesString(configuration.AVAILABLE_FILTERS, filter) {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidFilterName,
		)
	}

	partials := strings.Split(file.Filename, ".")
	fileExtension := strings.ToLower(partials[len(partials)-1])
	if fileExtension == "" ||
		!(fileExtension == "jpeg" ||
			fileExtension == "jpg" ||
			fileExtension == "png") {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidFileFormat,
		)
	}

	fileHandle, readerError := file.Open()
	if readerError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	format := fileExtension
	var processingError error
	var result io.Reader

	if filter == "binary" ||
		filter == "boxBlur" ||
		filter == "brightness" ||
		filter == "contrast" ||
		filter == "hueRotate" ||
		filter == "kuwahara" ||
		filter == "rotateFixed" ||
		filter == "sharpen" ||
		filter == "solarize" {
		threshold := context.FormValue("threshold")
		if threshold == "" {
			return fiber.NewError(
				fiber.StatusBadRequest,
				configuration.RESPONSE_MESSAGES.MissingThresholdValue,
			)
		}
		convertedThreshold, convertationError := strconv.Atoi(threshold)
		if convertationError != nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				configuration.RESPONSE_MESSAGES.InvalidThresholdValue,
			)
		}
		if filter == "binary" {
			result, format, processingError = brille.Binary(
				fileHandle,
				uint8(convertedThreshold),
			)
		}
		if filter == "boxBlur" {
			result, format, processingError = brille.BoxBlur(
				fileHandle,
				uint(utilities.Clamp(convertedThreshold, 25, 0)),
			)
		}
		if filter == "brightness" {
			result, format, processingError = brille.Brightness(
				fileHandle,
				convertedThreshold,
			)
		}
		if filter == "contrast" {
			result, format, processingError = brille.Contrast(
				fileHandle,
				convertedThreshold,
			)
		}
		if filter == "hueRotate" {
			result, format, processingError = brille.HueRotate(
				fileHandle,
				convertedThreshold,
			)
		}
		if filter == "kuwahara" {
			result, format, processingError = brille.Kuwahara(
				fileHandle,
				uint(utilities.Clamp(convertedThreshold, 25, 0)),
			)
		}
		if filter == "rotateFixed" {
			result, format, processingError = brille.RotateFixed(
				fileHandle,
				uint(convertedThreshold),
			)
		}
		if filter == "sharpen" {
			result, format, processingError = brille.Sharpen(
				fileHandle,
				uint(convertedThreshold),
			)
		}
		if filter == "solarize" {
			result, format, processingError = brille.Solarize(
				fileHandle,
				uint8(convertedThreshold),
			)
		}
	}

	if filter == "colorInversion" {
		result, format, processingError = brille.ColorInversion(fileHandle)
	}

	if filter == "eightColors" {
		result, format, processingError = brille.EightColors(fileHandle)
	}

	if filter == "emboss" {
		result, format, processingError = brille.Emboss(fileHandle)
	}

	if filter == "flip" {
		direction := context.FormValue("flipDirection")
		if direction == "" {
			return fiber.NewError(
				fiber.StatusBadRequest,
				configuration.RESPONSE_MESSAGES.MissingFlipDirection,
			)
		}
		result, format, processingError = brille.Flip(fileHandle, direction)
	}

	if filter == "gammaCorrection" {
		threshold := context.FormValue("threshold")
		if threshold == "" {
			return fiber.NewError(
				fiber.StatusBadRequest,
				configuration.RESPONSE_MESSAGES.MissingThresholdValue,
			)
		}
		convertedThreshold, convertationError := strconv.ParseFloat(threshold, 64)
		if convertationError != nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				configuration.RESPONSE_MESSAGES.InvalidThresholdValue,
			)
		}
		result, format, processingError = brille.GammaCorrection(
			fileHandle,
			convertedThreshold,
		)
	}

	if filter == "grayscale" {
		grayscaleType := context.FormValue("grayscaleType")
		if grayscaleType == "" {
			return fiber.NewError(
				fiber.StatusBadRequest,
				configuration.RESPONSE_MESSAGES.MissingGrayscaleType,
			)
		}
		result, format, processingError = brille.Grayscale(
			fileHandle,
			grayscaleType,
		)
	}

	if filter == "laplacian" {
		result, format, processingError = brille.Laplacian(fileHandle)
	}

	if filter == "sepia" {
		result, format, processingError = brille.Sepia(fileHandle)
	}

	if filter == "sobel" {
		result, format, processingError = brille.Sobel(fileHandle)
	}

	if processingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	context.Set("Content-Type", "image/"+format)
	return context.SendStream(result)
}
