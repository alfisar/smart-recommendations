package errorhandler

const (
	// response code for failed auth
	FailedAuth = "0001"

	// response code for failed validation
	FailedValidationToken = "0002"

	// response code for error from repository
	ErrorRepository = "0003"

	// response code for failed from hashing
	FailedHash = "0004"

	// response code for failed from controller
	FailedController = "0005"

	// response code for failed parsing data
	FailedParsingData = "0006"
)

type ErrorData struct {
	ResponseCode string
	Message      string
	Data         interface{}
}

func Errorifpanic(err error) {
	if err != nil {
		panic(err)
	}
}

func FailedGetToken(err error) ErrorData {
	result := ErrorData{
		ResponseCode: FailedAuth,
		Message:      err.Error(),
	}
	return result
}

func FailedvalidationToken(err error) ErrorData {
	result := ErrorData{
		ResponseCode: FailedValidationToken,
		Message:      err.Error(),
	}
	return result
}

func ErrorRepo(err error) ErrorData {
	result := ErrorData{
		ResponseCode: ErrorRepository,
		Message:      err.Error(),
	}
	return result
}

func ErrorHash(err error) ErrorData {
	result := ErrorData{
		ResponseCode: FailedHash,
		Message:      err.Error(),
	}
	return result
}

func ErrValidationService(responseCode string, err error) ErrorData {
	result := ErrorData{
		ResponseCode: responseCode,
		Message:      err.Error(),
	}
	return result
}

func ErrController(err error) ErrorData {
	result := ErrorData{
		ResponseCode: FailedController,
		Message:      err.Error(),
	}

	return result
}

func ErrParsingData(responseCode string, err error) ErrorData {
	result := ErrorData{
		ResponseCode: FailedParsingData,
		Message:      err.Error(),
	}
	return result
}
