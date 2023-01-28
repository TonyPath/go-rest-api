package http

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func JSONMiddleware(next http.Handler) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(h)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {

		log.WithFields(
			log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			},
		).Info("handling request")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(h)
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(h)
}
