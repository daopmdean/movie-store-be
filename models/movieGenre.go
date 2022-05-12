package models

import "time"

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   string    `json:"movieId"`
	GenreID   string    `json:"genreId"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
