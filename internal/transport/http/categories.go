package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/orzzet/ropero-solidario-api/src/validators"
	"net/http"
	"strconv"
)

func (h *Handler) createCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, validations := validators.CreateCategories(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	newCategories, err := h.Service.CreateCategories(data)
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(newCategories); err != nil {
		throwInternalError(w, err)
		return
	}
}

func (h *Handler) getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories, err := h.Service.GetCategories()
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		throwInternalError(w, err)
	}
}

func (h *Handler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["categoryId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid userId"))
	}
	err = h.Service.DeleteCategory(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
}
