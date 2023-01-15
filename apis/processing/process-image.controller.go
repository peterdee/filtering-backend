package processing

import (
	"io"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/brille"
	"github.com/julyskies/gohelpers"

	"filtering/configuration"
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

	if filter == "binary" {
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
		result, format, processingError = brille.Binary(
			fileHandle,
			uint(convertedThreshold),
		)
	}

	if filter == "boxBlur" {
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
		result, format, processingError = brille.BoxBlur(
			fileHandle,
			uint(convertedThreshold),
		)
	}

	if filter == "brightness" {
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
		result, format, processingError = brille.Brightness(
			fileHandle,
			convertedThreshold,
		)
	}

	if filter == "colorInversion" {
		result, format, processingError = brille.ColorInversion(fileHandle)
	}

	if filter == "contrast" {
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
		result, format, processingError = brille.Contrast(
			fileHandle,
			convertedThreshold,
		)
	}

	if filter == "eightColors" {
		result, format, processingError = brille.EightColors(fileHandle)
	}

	if filter == "emboss" {
		result, format, processingError = brille.EmbossFilter(fileHandle)
	}

	if filter == "flipHorizontal" {
		result, format, processingError = brille.FlipHorizontal(fileHandle)
	}

	if filter == "flipVertical" {
		result, format, processingError = brille.FlipVertical(fileHandle)
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
		if grayscaleType != brille.GRAYSCALE_AVERAGE &&
			grayscaleType != brille.GRAYSCALE_LUMINOCITY {
			return fiber.NewError(
				fiber.StatusBadRequest,
				configuration.RESPONSE_MESSAGES.InvalidGrayscaleType,
			)
		}
		result, format, processingError = brille.Grayscale(
			fileHandle,
			brille.GRAYSCALE_AVERAGE,
		)
	}

	if filter == "hueRotate" {
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
		result, format, processingError = brille.HueRotate(
			fileHandle,
			convertedThreshold,
		)
	}

	if filter == "laplasian" {
		result, format, processingError = brille.LaplasianFilter(fileHandle)
	}

	if filter == "rotate90" {
		result, format, processingError = brille.Rotate90(fileHandle)
	}

	if filter == "rotate180" {
		result, format, processingError = brille.Rotate180(fileHandle)
	}

	if filter == "rotate270" {
		result, format, processingError = brille.Rotate270(fileHandle)
	}

	if filter == "sepia" {
		result, format, processingError = brille.Sepia(fileHandle)
	}

	if filter == "sobel" {
		result, format, processingError = brille.SobelFilter(fileHandle)
	}

	if filter == "solarize" {
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
		result, format, processingError = brille.Solarize(
			fileHandle,
			uint(convertedThreshold),
		)
	}

	if processingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	context.Set("Content-Type", "image/"+format)
	return context.SendStream(result)
}
