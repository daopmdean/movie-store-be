package models

import "time"

type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"releaseDate"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaaRating"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	MovieGenre  []MovieGenre `json:"movieGenre"`
}
