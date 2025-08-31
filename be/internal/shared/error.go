package shared

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// implement error interface
func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// helper constructor
func WriteError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

// predefined error
var (
	ErrBadRequest          = WriteError(400, "Bad Request")
	ErrNotFound            = WriteError(404, "Not Found")
	ErrInternalServerError = WriteError(500, "Internal Server Error")
)
