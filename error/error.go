package error

import "net/http"

type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

func (e *Error) Error() string { return e.Message }

func New(status int, code, message string) *Error {
	return &Error{HTTPStatus: status, Code: code, Message: message}
}

var (
	ErrInternal     = New(http.StatusInternalServerError, "internal_error", "An unexpected internal error occurred.")
	ErrUnauthorized = New(http.StatusUnauthorized, "unauthorized", "Authentication failed or was not provided.")
)
