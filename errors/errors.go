package errors

import "fmt"

type ErrCode string

const (
	ErrCodeNotFound        ErrCode = "NOT_FOUND"
	ErrCodeAlreadyExists   ErrCode = "ALREADY_EXISTS"
	ErrCodeInvalidInput    ErrCode = "INVALID_INPUT"
	ErrCodeInternal        ErrCode = "INTERNAL_ERROR"
	ErrCodeUnauthorized    ErrCode = "UNAUTHORIZED"
	ErrCodeDatabaseFailure ErrCode = "DATABASE_FAILURE"
)

type Error interface {
	Error() string
	Unwrap() error
}

type ErrorImpl struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
	Err     error   `json:"-"`
}

func (e *ErrorImpl) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *ErrorImpl) Unwrap() error {
	return e.Err
}

func New(code ErrCode, message string, err error) *ErrorImpl {
	return &ErrorImpl{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
