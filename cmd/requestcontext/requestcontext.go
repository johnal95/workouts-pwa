package requestcontext

import (
	"context"
	"net/http"

	"github.com/johnal95/workouts-pwa/cmd/store"
)

type contextKey string

const userContextKey = contextKey("ctx.user")

func GetUser(r *http.Request) *store.User {
	user, ok := r.Context().Value(userContextKey).(*store.User)
	if !ok {
		return nil
	}
	return user
}

func SetUser(r *http.Request, user *store.User) *http.Request {
	return r.WithContext(
		context.WithValue(r.Context(), userContextKey, user),
	)
}
