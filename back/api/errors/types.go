package errors

import "errors"

type ErrorBody struct {
	GrpcCode int32       `json:"grpcCode"`
	Message  string      `json:"message"`
	Details  interface{} `json:"details"`
}

type ErrorDetail struct {
	Code    int32  `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

const ErrorDetailKey = "error-detail" // converted to lowdercase in setTrailer

var (
	Period              = ErrorDetail{Name: "period", Code: 100, Message: "Error: period length should be 6"}
	InvalidUserId       = ErrorDetail{Name: "userId", Code: 200, Message: "Error: userId is invalid"}
	InvalidPeriod       = ErrorDetail{Name: "period", Code: 201, Message: "Error: period is invalid"}
	InvalidDepartmentId = ErrorDetail{Name: "departmentId", Code: 202, Message: "Error: departmentId is invalid"}
	InvalidAimId        = ErrorDetail{Name: "aimId", Code: 203, Message: "Error: aimId is invalid"}
	InvalidAimNumber    = ErrorDetail{Name: "aimNumber", Code: 204, Message: "Error: aimNumber is invalid"}
	ErrNotFound         = errors.New("Error: Not Found")
	ErrRecordExist      = errors.New("Error: Record Exist")
	ErrNoContent        = errors.New("Error: No Content")

	NotEvenAToken          = errors.New("That's not even a token")
	TimingIsEverything     = errors.New("Timing is everything")
	CouldntHandleThisToken = errors.New("Couldn't handle this token")
)
