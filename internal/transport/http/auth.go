package http

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/orzzet/ropero-solidario-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var credentials models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		fmt.Fprintf(w, err.Error())
	}

	hashedPassword, err := h.UserService.GetUserHashedPassword(credentials.Email)

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
