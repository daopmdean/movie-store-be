package models

import "time"

type Genre struct {
	ID        string    `json:"id"`
	GenreName string    `json:"genreName"`
	CreateAt  time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
}
