package errors

func (e *AppError) GetCode() ErrCode   { return e.Code }
func (e *AppError) GetMessage() string { return e.Message }
func (e *AppError) GetDetails() any    { return e.Details }
func (e *AppError) Cause() error       { return e.Err }
func (e *AppError) WithDetails(details any) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: e.Message,
		Err:     e.Err,
		Details: details,
	}
}
