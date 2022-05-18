package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/daopmdean/movie-store-be/models"
	"github.com/julienschmidt/httprouter"
)

type jsonRes struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJson(w, err)
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.GetAll()
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJson(w, err)
	}
}

func (app *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreId, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	movies, err := app.models.DB.GetAll(genreId)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJson(w, err)
	}
}

type MoviePayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	ReleaseDate string `json:"releaseDate"`
	Runtime     string `json:"runtime"`
	Rating      string `json:"rating"`
	MPAARating  string `json:"mpaaRating"`
}

func (app *application) editMovie(w http.ResponseWriter, r *http.Request) {
	var moviePayload MoviePayload

	err := json.NewDecoder(r.Body).Decode(&moviePayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	var movie models.Movie
	movie.ID, _ = strconv.Atoi(moviePayload.ID)
	movie.Title = moviePayload.Title
	movie.Description = moviePayload.Description
	movie.ReleaseDate, err = time.Parse("2006-01-02", moviePayload.ReleaseDate)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, _ = strconv.Atoi(moviePayload.Runtime)
	movie.Rating, _ = strconv.Atoi(moviePayload.Rating)
	movie.MPAARating = moviePayload.MPAARating
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	if movie.ID == 0 {
		err = app.models.DB.InsertMovie(movie)
		if err != nil {
			app.errorJson(w, err)
			return
		}
	} else {
		m, err := app.models.DB.Get(movie.ID)
		if err != nil {
			app.errorJson(w, err)
			return
		}
		movie.CreatedAt = m.CreatedAt

		err = app.models.DB.UpdateMovie(movie)
		if err != nil {
			app.errorJson(w, err)
			return
		}
	}

	ok := jsonRes{
		Ok: true,
	}

	err = app.writeJson(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJson(w, err)
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchMovie(w http.ResponseWriter, r *http.Request) {

}
