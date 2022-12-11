package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
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

// Accepsts a pointer to move a movie struct, which should
// contain the data for the new record to add to the database
func (m *MovieModel) Insert(movie *Movie) error {
	query := `
    INSERT INTO movies (title, year, runtime, genres) 
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_at, version`

	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

// Fetch a specific record from the movies table.
func (m *MovieModel) Get(id int64) (*Movie, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
    SELECT id, created_at, title, year, runtime, genres, version
    FROM movies
    WHERE id = $1`

	var movie Movie

	err := m.DB.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &movie, nil
}

// Update a specific record in the movies table.
func (m *MovieModel) Update(movie *Movie) error {
	query := `
    UPDATE movies 
    SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
    WHERE id = $5
    RETURNING version`

	args := []any{
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres),
		movie.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&movie.Version)
}

// Delete a specific record from the movies table.
func (m *MovieModel) Delete(id int64) error {
	return nil
}
