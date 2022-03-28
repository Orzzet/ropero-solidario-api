package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/orzzet/ropero-solidario-api/src/services"
	"net/http"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Secret string
	Router *mux.Router
	*services.Service
}

type Response struct {
	Message string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(db *gorm.DB, secret string) *Handler {
	return &Handler{
		Secret:  secret,
		Service: services.NewService(db),
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/auth", h.Auth).Methods("POST")
	h.Router.HandleFunc("/users", h.CreateUser).Methods("POST")
	h.Router.HandleFunc("/users/{userId}/approve", h.ApproveUser).Methods("POST")
	h.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(Response{Message: "Online"}); err != nil {
			panic(err)
		}
	})
}

func (h *Handler) ThrowError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	validationErrors := make(map[string]string)
	for _, valErr := range err.(validator.ValidationErrors) {
		validationErrors[valErr.Field()] = valErr.Tag()
	}
	formattedValidationErrors, _ := json.Marshal(validationErrors)
	w.Write(formattedValidationErrors)
}
