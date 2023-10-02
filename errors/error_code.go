package errors

type ErrorCode string

const (
	ErrorUserNotAcceptTermsAndConditions      ErrorCode = "0000001"
	ErrorInvalidEmail                         ErrorCode = "0000002"
	ErrorPasswordDoesntHaveTheRequestedFormat ErrorCode = "0000003"
	ErrorPasswordDoesntMatch                  ErrorCode = "0000004"
)
