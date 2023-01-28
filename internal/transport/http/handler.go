package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router        *mux.Router
	CommentSevice CommentService
	Server        *http.Server
}

func NewHandler(commentService CommentService) *Handler {
	h := Handler{
		CommentSevice: commentService,
		Router:        mux.NewRouter(),
	}

	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
	h.Router.Use(TimeoutMiddleware)

	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
	}

	return &h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "alice")
	})

	h.Router.HandleFunc("/api/v1/comments", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/v1/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comments/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comments/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	h.Server.Shutdown(ctx)
	log.Println("shutting down gracefully")

	return nil
}
