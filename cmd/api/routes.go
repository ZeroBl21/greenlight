package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Encapsulates the API routes of the application.
func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Error Routes
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Movies
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.listMoviesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)

	// Users
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	// Auth
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication",
		app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
