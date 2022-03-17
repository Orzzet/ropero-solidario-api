package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/orzzet/ropero-solidario-api/internal/models"
	"net/http"
	"strconv"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *models.Service
}

type Response struct {
	Message string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *models.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comments", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comments/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comments/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})
}

// GetComment gets a comment
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "%+v", comment)
}

// GetAllComments gets all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "%+v", comments)
}

// PostComment posts a comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// UpdateComment posts a comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	comment := models.Comment{
		Slug:   vars["slug"],
		Body:   vars["body"],
		Author: vars["author"],
	}
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	updatedComment, err := h.Service.UpdateComment(uint(id), comment)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "%+v", updatedComment)
}

// DeleteComment posts a comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	err = h.Service.DeleteComment(uint(id))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
