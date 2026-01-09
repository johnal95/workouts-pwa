package httpx

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type PublicError struct {
	StatusCode int
	Message    string
	Details    any
	Err        error
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

func NotFound(err error, msg string, details any) PublicError {
	return PublicError{
		StatusCode: http.StatusNotFound,
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
