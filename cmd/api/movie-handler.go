package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println("error dude")
		app.errorJson(w, err)
		return
	}

	app.writeJson(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {

}
