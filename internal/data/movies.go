package data

import (
	"database/sql"
	"time"

	"github.com/zerobl21/greenlight/internal/validator"
)

// Holds the Movies data of the application.
type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	CreatedAt time.Time `json:"-"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	// Title Validation
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	// Year Validation
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	// Runtime Validation
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	// Genres Validation
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicated values")
}

// Wraps the sql.DB connection pool.
type MovieModel struct {
	DB *sql.DB
}

// Insert a new record in the movie table.
func (m *MovieModel) Insert(movie *Movie) error {
	return nil
}

// Fetch a specific record from the movies table.
func (m *MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

// Update a specific record in the movies table.
func (m *MovieModel) Update(movie *Movie) error {
	return nil
}

// Delete a specific record from the movies table.
func (m *MovieModel) Delete(id int64) error {
	return nil
}
