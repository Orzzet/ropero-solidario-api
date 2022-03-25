package http

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/orzzet/ropero-solidario-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Secret  string
	Router  *mux.Router
	Service *models.Service
}

type Response struct {
	Message string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *models.Service, secret string) *Handler {
	return &Handler{
		Service: service,
		Secret:  secret,
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

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var credentials models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		fmt.Fprintf(w, err.Error())
	}

	hashedPassword, err := h.Service.GetUserHashedPassword(credentials.Email)

	if err != nil {
		// If there is an issue with the database, return a 500 error
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":        credentials.Email,
		"creationDate": time.Now().Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(h.Secret))
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"token": tokenString}); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.IsApproved = false
	user.Password = string(hashedPassword)

	user, err = h.Service.CreateUser(user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func (h *Handler) ApproveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	user, err = h.Service.ApproveUser(uint(id))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
