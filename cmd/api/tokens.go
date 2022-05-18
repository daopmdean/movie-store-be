package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/daopmdean/movie-store-be/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var validUser = models.User{
	ID:       175,
	Email:    "test@gmail.com",
	Password: "$2a$12$q7ZdXR7PRNN0RLaMaIBxUuhlFDMa5DmeRmZzV8WAekRaIcFjxI2U6", // hash of password
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	hashedPassword := validUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	var claims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		Issuer:    "authen-jwt",
		Audience:  "phamminhdao.com",
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := t.SignedString([]byte(app.config.jwt.secret))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, signedToken, "token")
	if err != nil {
		app.errorJson(w, err)
	}
}
