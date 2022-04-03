package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/orzzet/ropero-solidario-api/src/models"
	"github.com/orzzet/ropero-solidario-api/src/validators"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, validations := validators.CreateUser(r)
	if validations != nil {
		json.NewEncoder(w).Encode(validations)
		return
	}

	newUser, err := h.Service.CreateUser(data)
	if err != nil {
		h.ThrowError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User

	users, err := h.Service.GetUsers()
	if err != nil {
		h.ThrowError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func (h *Handler) ApproveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	newUser, err := h.Service.ApproveUser(uint(id))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
