package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	usermodel "github.com/orzzet/ropero-solidario-api/src/user"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user usermodel.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.IsApproved = false
	user.Password = string(hashedPassword)

	user, err = h.UserService.CreateUser(user)
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
	var user usermodel.User
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	user, err = h.UserService.ApproveUser(uint(id))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
