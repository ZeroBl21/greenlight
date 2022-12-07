package data

import "time"

// Holds the Movies data of the application.
type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty,string"`
	Genres    []string  `json:"genres,omitempty"`
	CreatedAt time.Time `json:"-"`
	Version   int32     `json:"version"`
}
