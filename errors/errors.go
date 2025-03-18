package errors

import (
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
	ErrCodeUnauthorized     ErrCode = "UNAUTHORIZED"
	ErrCodePermissionDenied ErrCode = "PERMISSION_DENIED"
)

type ErrorImpl struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
	Err     error   `json:"-"`
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

func (e *ErrorImpl) GRPCStatus() *status.Status {
	return status.New(e.toGrpcCode(), e.Message)
}

func (e *ErrorImpl) toGrpcCode() codes.Code {
	switch e.Code {
	case ErrCodeNotFound:
		return codes.NotFound
	case ErrCodeAlreadyExists:
		return codes.AlreadyExists
	case ErrCodeInvalidInput:
		return codes.InvalidArgument
	case ErrCodeUnauthorized:
		return codes.Unauthenticated
	case ErrCodePermissionDenied:
		return codes.PermissionDenied
	default:
		return codes.Internal
	}
}
