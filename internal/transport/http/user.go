package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/orzzet/ropero-solidario-api/src/models"
	"github.com/orzzet/ropero-solidario-api/src/validators"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, validations := validators.CreateUser(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}

	newUser, err := h.Service.CreateUser(data)
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		throwInternalError(w, err)
		return
	}
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User

	users, err := h.Service.GetUsers()
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		throwInternalError(w, err)
	}
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid userId"))
	}
	user, err := h.Service.GetUser(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		throwInternalError(w, err)
		return
	}
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid userId"))
	}
	err = h.Service.DeleteUser(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
}

func (h *Handler) approveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid userId"))
	}
	user, err := h.Service.ApproveUser(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		throwInternalError(w, err)
		return
	}
}

func (h *Handler) resetUserPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid userId"))
	}
	data, validations := validators.ResetPassword(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	password := data["password"].(string)
	user, err := h.Service.ResetPassword(uint(id), password)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		throwInternalError(w, err)
		return
	}
}
