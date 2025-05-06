package errors

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrCode string

const (
	ErrCodeNotFound         ErrCode = "NOT_FOUND"
	ErrCodeAlreadyExists    ErrCode = "ALREADY_EXISTS"
	ErrCodeInvalidInput     ErrCode = "INVALID_INPUT"
	ErrCodeInternal         ErrCode = "INTERNAL_ERROR"
	ErrCodeInvalidArgument  ErrCode = "INVALID_ARGUMENT"
	ErrCodeUnauthorized     ErrCode = "UNAUTHORIZED"
	ErrCodePermissionDenied ErrCode = "PERMISSION_DENIED"
)

type ErrorImpl struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
	Err     error   `json:"-"`
	Details any     `json:"details,omitempty"`
}

func New(code ErrCode, message string, err error) *ErrorImpl {
	return &ErrorImpl{
		Code:    code,
		Message: message,
		Err:     err,
	}
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

func (e *ErrorImpl) WithDetails(details any) *ErrorImpl {
	e.Details = details
	return e
}

func (e *ErrorImpl) GRPCStatus() *status.Status {
	return status.New(e.toGRPCCode(), e.Message)
}

func (e *ErrorImpl) toGRPCCode() codes.Code {
	switch e.Code {
	case ErrCodeNotFound:
		return codes.NotFound
	case ErrCodeAlreadyExists:
		return codes.AlreadyExists
	case ErrCodeInvalidInput, ErrCodeInvalidArgument:
		return codes.InvalidArgument
	case ErrCodeUnauthorized:
		return codes.Unauthenticated
	case ErrCodePermissionDenied:
		return codes.PermissionDenied
	default:
		return codes.Internal
	}
}

func (e *ErrorImpl) Is(target error) bool {
	t, ok := target.(*ErrorImpl)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func AsErrCode(err error) (ErrCode, bool) {
	var e *ErrorImpl
	if errors.As(err, &e) {
		return e.Code, true
	}
	return "", false
}
