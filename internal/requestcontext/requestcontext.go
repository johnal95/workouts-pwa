package requestcontext

import (
	"context"
	"net/http"
)

type contextKey string

const userIDContextKey = contextKey("ctx.user_id")
const requestIDContextKey = contextKey("ctx.request_id")

func UserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(userIDContextKey).(string)
	return userID, ok
}

func MustUserID(r *http.Request) string {
	userID, ok := r.Context().Value(userIDContextKey).(string)
	if !ok {
		panic("requestcontext: user ID missing")
	}
	return userID
}

func SetUserID(r *http.Request, userID string) *http.Request {
	return r.WithContext(
		context.WithValue(r.Context(), userIDContextKey, userID),
	)
}

func RequestID(r *http.Request) (string, bool) {
	requestID, ok := r.Context().Value(requestIDContextKey).(string)
	return requestID, ok
}

func MustRequestID(r *http.Request) string {
	requestID, ok := r.Context().Value(requestIDContextKey).(string)
	if !ok {
		panic("requestcontext: request ID missing")
	}
	return requestID
}

func SetRequestID(r *http.Request, requestID string) *http.Request {
	return r.WithContext(
		context.WithValue(r.Context(), requestIDContextKey, requestID),
	)
}
