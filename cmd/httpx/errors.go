package httpx

import "net/http"

type PublicError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e PublicError) Error() string {
	return e.Err.Error()
}

func InvalidRequestBody(err error) PublicError {
	return BadRequest("invalid request body", err)
}

func BadRequest(msg string, err error) PublicError {
	return PublicError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Err:        err,
	}
}

func NotFound(msg string, err error) PublicError {
	return PublicError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
		Err:        err,
	}
}

func Conflict(msg string, err error) PublicError {
	return PublicError{
		StatusCode: http.StatusConflict,
		Message:    msg,
		Err:        err,
	}
}

func Internal(err error) PublicError {
	return PublicError{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
		Err:        err,
	}
}
