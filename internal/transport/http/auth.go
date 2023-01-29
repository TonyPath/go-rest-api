package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type apiHandlerFunc func(w http.ResponseWriter, r *http.Request)

func JWTAuth(next apiHandlerFunc) apiHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if !validateToken(authHeaderParts[1]) {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}

}

func validateToken(token string) bool {
	var mySigningKey = []byte("secret")

	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return tkn.Valid

}
