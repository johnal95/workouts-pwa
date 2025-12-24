package requestcontext

import (
	"context"
	"net/http"
)

type contextKey string

const userIDContextKey = contextKey("ctx.user_id")
const requestIDContextKey = contextKey("ctx.request_id")

func GetUserID(r *http.Request) *string {
	userID, ok := r.Context().Value(userIDContextKey).(*string)
	if !ok {
		return nil
	}
	return userID
}

func SetUserID(r *http.Request, userID string) *http.Request {
	return r.WithContext(
		context.WithValue(r.Context(), userIDContextKey, &userID),
	)
}

func GetRequestID(r *http.Request) *string {
	requestID, ok := r.Context().Value(requestIDContextKey).(*string)
	if !ok {
		return nil
	}
	return requestID
}

func SetRequestID(r *http.Request, requestID string) *http.Request {
	return r.WithContext(
		context.WithValue(r.Context(), requestIDContextKey, &requestID),
	)
}
