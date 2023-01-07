package configuration

// this should corellate with julyskies/brille
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
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidData:         "INVALID_DATA",
	InvalidFileFormat:   "INVALID_FILE_FORMAT",
	InvalidFilterName:   "INVALID_FILTER_NAME",
	MissingData:         "MISSING_DATA",
	OK:                  "OK",
	TooManyRequests:     "TOO_MANY_REQUESTS",
}
