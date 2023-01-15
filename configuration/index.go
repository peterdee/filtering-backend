package configuration

// this should correlate with julyskies/brille
var AVAILABLE_FILTERS = []string{
	"binary",
	"boxBlur",
	"brightness",
	"colorInversion",
	"contrast",
	"eightColors",
	"emboss",
	"flipHorizontal",
	"flipVertical",
	"gammaCorrection",
	"grayscale",
	"hueRotate",
	"laplasian",
	"rotate90",
	"rotate180",
	"rotate270",
	"sepia",
	"sobel",
	"solarize",
}

const DEFAULT_ALLOWED_ORIGINS string = "http://localhost:3000"

const DEFAULT_PORT string = "8910"

var RESPONSE_MESSAGES = ResponseMessagesStruct{
	InternalServerError:   "INTERNAL_SERVER_ERROR",
	InvalidFileFormat:     "INVALID_FILE_FORMAT",
	InvalidFilterName:     "INVALID_FILTER_NAME",
	InvalidGrayscaleType:  "INVALID_GRAYSCALE_TYPE",
	InvalidThresholdValue: "INVALID_THRESHLOD_VALUE",
	MissingFilterName:     "MISSING_FILTER_NAME",
	MissingGrayscaleType:  "MISSING_GRAYSCALE_TYPE",
	MissingThresholdValue: "MISSING_THRESHOLD_VALUE",
	OK:                    "OK",
	TooManyRequests:       "TOO_MANY_REQUESTS",
}
