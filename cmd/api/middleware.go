package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		next.ServeHTTP(w, r)
	})
}

func (app *application) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// anonymous user
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			app.errorJson(w, fmt.Errorf("Invalid token"))
			return
		}

		if headerParts[0] != "Bearer" {
			app.errorJson(w, fmt.Errorf("Invalid token: %s", "Bearer not found"))
			return
		}

		tokenString := headerParts[1]
		var claims jwt.StandardClaims

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("Invalid Alg")
			}
			return []byte(app.config.jwt.secret), nil
		})
		if err != nil {
			app.errorJson(w, fmt.Errorf("Invalid token: %s", "Error when parsing"), http.StatusForbidden)
			return
		}

		if !token.Valid {
			app.errorJson(w, fmt.Errorf("Invalid token"), http.StatusForbidden)
			return
		}

		if claims.Audience != "phamminhdao.com" {
			app.errorJson(w, fmt.Errorf("Invalid token: %s", "Invalid Audience"), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
