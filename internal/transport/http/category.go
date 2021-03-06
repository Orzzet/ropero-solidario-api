package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) createCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, validations := h.Validator.CreateCategories(r)
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

func (h *Handler) createCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, validations := h.Validator.CreateCategory(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	newCategory, err := h.Service.CreateCategory(data)
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(newCategory); err != nil {
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

func (h *Handler) getCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["categoryId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid categoryId"))
	}
	category, err := h.Service.GetCategory(uint(id))
	if err != nil {
		throwNotFoundError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(category); err != nil {
		throwInternalError(w, err)
	}
}

func (h *Handler) editCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["categoryId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid categoryId"))
	}
	data, validations := h.Validator.CreateCategory(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	category, err := h.Service.EditCategory(uint(id), data)
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(category); err != nil {
		throwInternalError(w, err)
		return
	}
}

func (h *Handler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["categoryId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid categoryId"))
	}
	validations := h.Validator.DeleteCategory(uint(id))
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	err = h.Service.DeleteCategory(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Category not found"))
		return
	}
}
