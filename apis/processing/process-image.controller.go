package processing

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"strings"

	"github.com/gofiber/fiber/v2"

	"filtering/configuration"
)

// TODO: remove all of this since it has been moved to julyskies/brille

const GRAY_CALCULATION_AVERAGE string = "average"

const GRAY_CALCULATION_LUMINOCITY string = "luminocity"

func prepareImage(file io.Reader) ([][]color.Color, string, error) {
	content, format, decodingError := image.Decode(file)
	if decodingError != nil {
		return nil, "", decodingError
	}

	rect := content.Bounds()
	height, width := rect.Dy(), rect.Dx()
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), content, rect.Min, draw.Src)

	grid := make([][]color.Color, width)
	for x := 0; x < width; x += 1 {
		col := make([]color.Color, height)
		for y := 0; y < height; y += 1 {
			col[y] = rgba.At(x, y)
		}
		grid[x] = col
	}

	return grid, format, nil
}

func createGrid(width, height int) [][]color.Color {
	gridCopy := make([][]color.Color, width)
	for i := range gridCopy {
		gridCopy[i] = make([]color.Color, height)
	}
	return gridCopy
}

func gray(pixel color.Color, calculationType string) (gray uint8, alpha uint8) {
	R, G, B, A := pixel.RGBA()
	alpha = uint8(A)
	if calculationType == GRAY_CALCULATION_LUMINOCITY {
		gray = uint8(
			math.Round(
				(float64(uint8(R))*0.21 + float64(uint8(G))*0.72 + float64(uint8(B))*0.07),
			),
		)
		return
	}
	gray = uint8(
		math.Round(
			(float64(uint8(R)) + float64(uint8(G)) + float64(uint8(B))) / 3.0,
		),
	)
	return
}

func encodeResult(result [][]color.Color, format string) (io.Reader, error) {
	width, height := len(result), len(result[0])
	resultingImage := image.NewNRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			resultingImage.Set(x, y, result[x][y])
		}
	}

	var buffer bytes.Buffer
	resultingFile := io.Writer(&buffer)
	if format == "png" {
		encodingError := png.Encode(
			resultingFile,
			resultingImage.SubImage(resultingImage.Rect),
		)
		if encodingError != nil {
			return nil, encodingError
		}
	} else {
		encodingError := jpeg.Encode(
			resultingFile,
			resultingImage.SubImage(resultingImage.Rect),
			nil,
		)
		if encodingError != nil {
			return nil, encodingError
		}
	}
	return bytes.NewReader(buffer.Bytes()), nil
}

func grayscale(source [][]color.Color, grayscaleType string) [][]color.Color {
	width := len(source)
	height := len(source[0])

	destination := createGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			var grayColor, alpha uint8
			if grayscaleType == GRAY_CALCULATION_LUMINOCITY {
				grayColor, alpha = gray(source[x][y], GRAY_CALCULATION_LUMINOCITY)
			} else {
				grayColor, alpha = gray(source[x][y], GRAY_CALCULATION_AVERAGE)
			}
			destination[x][y] = color.RGBA{grayColor, grayColor, grayColor, alpha}
		}
	}
	return destination
}

func processImageController(context *fiber.Ctx) error {
	// get file
	file, fileError := context.FormFile("image")
	// TODO: check filter (depends on julyskies/brille module)
	// filter := context.FormValue("filter")
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

	preparedImage, format, preparationError := prepareImage(fileHandle)
	if preparationError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	grayscaled := grayscale(preparedImage, GRAY_CALCULATION_LUMINOCITY)
	encoded, encodingError := encodeResult(grayscaled, format)
	if encodingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	context.Set("Content-Type", "image/"+format)
	// TODO: set filename if necessary
	// context.Set("Content-Disposition", "attachment; filename="filename.jpg"")
	return context.SendStream(encoded)
}
