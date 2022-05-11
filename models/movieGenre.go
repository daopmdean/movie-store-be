package models

import "time"

type MovieGenre struct {
	ID       string    `json:"id"`
	MovieID  string    `json:"movieId"`
	GenreID  string    `json:"genreId"`
	Genre    Genre     `json:"genre"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}
