package errors

import (
	"net/http"
	"sync"
)

var (
	httpMapping sync.Map
)

func init() {
	defaults := map[ErrCode]int{
		ErrCodeNotFound:         http.StatusNotFound,
		ErrCodeAlreadyExists:    http.StatusConflict,
		ErrCodeInvalidInput:     http.StatusBadRequest,
		ErrCodeInvalidArgument:  http.StatusBadRequest,
		ErrCodeInternal:         http.StatusInternalServerError,
		ErrCodeUnauthorized:     http.StatusUnauthorized,
		ErrCodePermissionDenied: http.StatusForbidden,
	}
	for k, v := range defaults {
		httpMapping.Store(k, v)
	}
}

func (e *AppError) HTTPStatus() int {
	if val, ok := httpMapping.Load(e.Code); ok {
		return val.(int)
	}
	return http.StatusInternalServerError
}

func FromHTTPStatus(code int, message string) *AppError {
	return New(httpToErrorCode(code), message, nil)
}

func httpToErrorCode(code int) ErrCode {
	switch code {
	case http.StatusNotFound:
		return ErrCodeNotFound
	case http.StatusConflict:
		return ErrCodeAlreadyExists
	case http.StatusBadRequest:
		return ErrCodeInvalidArgument
	case http.StatusUnauthorized:
		return ErrCodeUnauthorized
	case http.StatusForbidden:
		return ErrCodePermissionDenied
	default:
		return ErrCodeInternal
	}
}

func RegisterHTTPMapping(code ErrCode, statusCode int) {
	httpMapping.Store(code, statusCode)
}
