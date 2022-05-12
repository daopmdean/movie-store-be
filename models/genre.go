package models

import "time"

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genreName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
