package main

import (
	"fmt"
	"net/http"
)

// Creates a new movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movie")
}

// Returns the movie with the especified id
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
