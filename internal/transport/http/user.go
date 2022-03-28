package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/orzzet/ropero-solidario-api/src/models"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Fprintf(w, err.Error())
	}

	newUser, err := h.Service.CreateUser(user)
	if err != nil {
		h.ThrowError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
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
