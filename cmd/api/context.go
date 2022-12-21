package main

import (
	"context"
	"net/http"

	"github.com/zerobl21/greenlight/internal/data"
)

// Custom type, with the underlynig type string.
type contextKey string

const userContextKey = contextKey("user")

// Returns a new copy of the request with the provided User struct added to the context.
func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)

	return r.WithContext(ctx)
}

// Retrieves the User struct form the request context.
func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
