package models

import "time"

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movieId"`
	GenreID   int       `json:"genreId"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
