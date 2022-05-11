package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/daopmdean/movie-store-be/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	movie := models.Movie{
		ID:          id,
		Title:       "Movie",
		Description: "Des",
		Year:        2121,
		ReleaseDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-12",
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}

	app.writeJson(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {

}
