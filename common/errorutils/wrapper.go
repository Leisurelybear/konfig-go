package errorutils

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
