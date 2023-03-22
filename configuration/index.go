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
	"flip",
	"gammaCorrection",
	"gaussianBlur",
	"grayscale",
	"hueRotate",
	"kuwahara",
	"laplacian",
	"rotateFixed",
	"sepia",
	"sharpen",
	"sobel",
	"solarize",
}

const DEFAULT_ALLOWED_ORIGINS string = "http://localhost:3000"

const DEFAULT_PORT string = "8910"

var RESPONSE_MESSAGES = ResponseMessagesStruct{
	InternalServerError:   "INTERNAL_SERVER_ERROR",
	InvalidFileFormat:     "INVALID_FILE_FORMAT",
	InvalidFilterName:     "INVALID_FILTER_NAME",
	InvalidThresholdValue: "INVALID_THRESHLOD_VALUE",
	MissingFilterName:     "MISSING_FILTER_NAME",
	MissingFlipDirection:  "MISSING_FLIP_DIRECTION",
	MissingGrayscaleType:  "MISSING_GRAYSCALE_TYPE",
	MissingThresholdValue: "MISSING_THRESHOLD_VALUE",
	OK:                    "OK",
	TooManyRequests:       "TOO_MANY_REQUESTS",
}
