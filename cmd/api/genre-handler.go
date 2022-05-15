package main

import "net/http"

func (app *application) getGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.AllGenres()
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}
