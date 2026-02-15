package httpx

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type PublicError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    any    `json:"details"`
	Err        error  `json:"-"`
}

func (e PublicError) Error() string {
	return e.Err.Error()
}

func ValidationError(err error, target any) error {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return InternalServerError(err)
	}
	fields := map[string]string{}
	for _, fe := range ve {
		fields[fe.Field()] = fe.Tag()
	}
	return BadRequest(err, "invalid request body", fields)

}

func BadRequest(err error, msg string, details any) PublicError {
	return PublicError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Details:    details,
		Err:        err,
	}
}

func Unauthorized(err error, msg string, details any) PublicError {
	return PublicError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
		Details:    details,
		Err:        err,
	}
}

func NotFound(err error, msg string, details any) PublicError {
	return PublicError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
		Details:    details,
		Err:        err,
	}
}

func NotImplemented(err error, msg string, details any) PublicError {
	return PublicError{
		StatusCode: http.StatusNotImplemented,
		Message:    msg,
		Details:    details,
		Err:        err,
	}
}

func Conflict(err error, msg string, details any) PublicError {
	return PublicError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Details:    details,
		Err:        err,
	}
}

func InternalServerError(err error) PublicError {
	return PublicError{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
		Err:        err,
	}
}
