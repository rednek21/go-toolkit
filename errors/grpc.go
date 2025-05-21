package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *AppError) GRPCStatus() *status.Status {
	return status.New(toGRPCCode(e.Code), e.Message)
}

func toGRPCCode(code ErrCode) codes.Code {
	switch code {
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
	case ErrCodeInternal:
		return codes.Internal
	}

	return codes.Unknown
}

func FromGRPC(st *status.Status) *AppError {
	return New(fromGRPCCode(st.Code()), st.Message(), nil)
}

func fromGRPCCode(c codes.Code) ErrCode {
	switch c {
	case codes.NotFound:
		return ErrCodeNotFound
	case codes.AlreadyExists:
		return ErrCodeAlreadyExists
	case codes.InvalidArgument:
		return ErrCodeInvalidArgument
	case codes.Unauthenticated:
		return ErrCodeUnauthorized
	case codes.PermissionDenied:
		return ErrCodePermissionDenied
	default:
		return ErrCodeInternal
	}
}
