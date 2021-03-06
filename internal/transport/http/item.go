package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, validations := h.Validator.CreateItem(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	item, err := h.Service.CreateItem(data)
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		throwInternalError(w, err)
		return
	}
}

func (h *Handler) getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := h.Service.GetItems()
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(items); err != nil {
		throwInternalError(w, err)
	}
}

func (h *Handler) getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["itemId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid itemId"))
	}
	item, err := h.Service.GetItem(uint(id))
	if err != nil {
		throwNotFoundError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		throwInternalError(w, err)
	}
}

func (h *Handler) editItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["itemId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid itemId"))
	}
	data, validations := h.Validator.CreateItem(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	item, err := h.Service.EditItem(uint(id), data)
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		throwInternalError(w, err)
		return
	}
}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["itemId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid itemId"))
	}
	validations := h.Validator.DeleteItem(uint(id))
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	err = h.Service.DeleteItem(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Item not found"))
		return
	}
}
