package httpx

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Details any    `json:"details,omitempty"`
}

func RespondJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set(HeaderContentType, "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func RespondError(w http.ResponseWriter, err error) {
	var pe PublicError

	if errors.As(err, &pe) {
		RespondJSON(w, pe.StatusCode, ErrorResponse{
			Error:   pe.Message,
			Details: pe.Details,
		})
		return
	}

	RespondJSON(w, http.StatusInternalServerError, InternalServerError(err))
}

func RespondNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
