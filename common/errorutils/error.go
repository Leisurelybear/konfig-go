package errorutils

import "fmt"

const (
	Success = iota
	ServerErr
	NotFound
)

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d|%s", e.Code, e.Message)
}
