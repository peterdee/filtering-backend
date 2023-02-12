package configuration

type ResponseMessagesStruct struct {
	InternalServerError   string
	InvalidFileFormat     string
	InvalidFilterName     string
	InvalidThresholdValue string
	MissingFilterName     string
	MissingFlipDirection  string
	MissingGrayscaleType  string
	MissingThresholdValue string
	OK                    string
	TooManyRequests       string
}
