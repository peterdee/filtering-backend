package configuration

type ResponseMessagesStruct struct {
	InternalServerError   string
	InvalidFileFormat     string
	InvalidFilterName     string
	InvalidGrayscaleType  string
	InvalidThresholdValue string
	MissingFilterName     string
	MissingGrayscaleType  string
	MissingThresholdValue string
	OK                    string
	TooManyRequests       string
}
