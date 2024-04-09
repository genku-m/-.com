package errors

type ErrCode string

var (
	ErrNotFound        ErrCode = "Not found"
	ErrInvalidArgument ErrCode = "Invalid argument"
	ErrInternal        ErrCode = "Internal server error"
	ErrUnauthorized    ErrCode = "Unauthorized"
)

type ServerError struct {
	Msg     ErrorMsg
	ErrCode ErrCode
	Err     error
}

type ErrorMsg string

var (
	ErrNotFoundMsg        ErrorMsg = "Invoice not found"
	ErrInvalidArgumentMsg ErrorMsg = "Invalid argument"
	ErrInternalMsg        ErrorMsg = "Internal server error"
	ErrUnauthorizedMsg    ErrorMsg = "Unauthorized"
)

func NewNotFoundError(err error) *ServerError {
	return &ServerError{
		Msg:     ErrNotFoundMsg,
		ErrCode: ErrNotFound,
		Err:     err,
	}
}

func NewInvalidArgumentError(err error) *ServerError {
	return &ServerError{
		Msg:     ErrInvalidArgumentMsg,
		ErrCode: ErrInvalidArgument,
		Err:     err,
	}
}

func NewInternalError(err error) *ServerError {
	return &ServerError{
		Msg:     ErrInternalMsg,
		ErrCode: ErrInternal,
		Err:     err,
	}
}

func NewUnauthorizedError(err error) *ServerError {
	return &ServerError{
		Msg:     ErrUnauthorizedMsg,
		ErrCode: ErrUnauthorized,
		Err:     err,
	}
}

func (s *ServerError) Error() string {
	return s.Err.Error()
}

func IsErrNotFound(err error) bool {
	serverError, ok := err.(*ServerError)
	if !ok {
		return false
	}
	return serverError.ErrCode == ErrNotFound
}

func IsErrInvalidArgument(err error) bool {
	serverError, ok := err.(*ServerError)
	if !ok {
		return false
	}
	return serverError.ErrCode == ErrInvalidArgument
}

func IsErrInternal(err error) bool {
	serverError, ok := err.(*ServerError)
	if !ok {
		return false
	}
	return serverError.ErrCode == ErrInternal
}

func IsErrUnauthrized(err error) bool {
	serverError, ok := err.(*ServerError)
	if !ok {
		return false
	}
	return serverError.ErrCode == ErrUnauthorized
}
