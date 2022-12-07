package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Encapsulates the API routes of the application.
func (app *application) routes() *httprouter.Router {
  router := httprouter.New()

  router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

  // Movies
  router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
  router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

  return router
}
