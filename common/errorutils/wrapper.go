package errorutils

import "errors"

var (
	ErrInvalidParameters = errors.New("parameter is invalid")
	ErrInvalidDuplicated = errors.New("item is duplicated")
	ErrNotFound          = errors.New("item not found")
)

func Wrap(err error) *Error {
	if err == nil {
		return &Error{
			Code: Success,
		}
	}
	wrappedErr := &Error{
		Code:    ServerErr,
		Message: err.Error(),
	}
	return wrappedErr
}
