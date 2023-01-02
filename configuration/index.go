package configuration

const DEFAULT_PORT string = "8910"

var RESPONSE_MESSAGES = ResponseMessagesStruct{
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidData:         "INVALID_DATA",
	MissingData:         "MISSING_DATA",
	OK:                  "OK",
	TooManyRequests:     "TOO_MANY_REQUESTS",
}
