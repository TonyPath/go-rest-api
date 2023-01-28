package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TonyPath/go-rest-api/internal/comment"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type CommentService interface {
	GetComment(ctx context.Context, id string) (comment.Comment, error)
	CreateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, uComment comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, id string) error
}

type Response struct {
	Message string
}

type newCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func (newCmtReq newCommentRequest) toComment() comment.Comment {
	return comment.Comment{
		Slug:   newCmtReq.Slug,
		Author: newCmtReq.Author,
		Body:   newCmtReq.Body,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var newCommentReq newCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&newCommentReq); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(newCommentReq)
	if err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
		return
	}

	cmt, err := h.CommentSevice.CreateComment(r.Context(), newCommentReq.toComment())
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		log.Println(err)
		return
	}

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.CommentSevice.GetComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.CommentSevice.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.CommentSevice.DeleteComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(Response{Message: "successfully deleted"})
	if err != nil {
		panic(err)
	}
}
