package errors

import (
	"errors"
	"fmt"
)

type ErrCode string

type AppError struct {
	Code    ErrCode
	Message string
	Details any
	Err     error
}

var _ error = (*AppError)(nil)

func New(code ErrCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *AppError) Error() string {
	msg := fmt.Sprintf("[%s] %s", e.Code, e.Message)
	if e.Details != nil {
		msg += fmt.Sprintf(" | details: %v", e.Details)
	}
	if e.Err != nil {
		msg += fmt.Sprintf(": %v", e.Err)
	}
	return msg
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func AsErrCode(err error) (ErrCode, bool) {
	var e *AppError
	if errors.As(err, &e) {
		return e.Code, true
	}
	return "", false
}
