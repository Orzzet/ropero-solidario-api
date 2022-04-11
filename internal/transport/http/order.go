package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/orzzet/ropero-solidario-api/src/validators"
	"net/http"
	"strconv"
)

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, validations := validators.CreateOrder(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	item, err := h.Service.CreateOrder(data)
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		throwInternalError(w, err)
		return
	}
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders, err := h.Service.GetOrders()
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		throwInternalError(w, err)
	}
}

func (h *Handler) getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	order, err := h.Service.GetOrder()
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(order); err != nil {
		throwInternalError(w, err)
	}
}

func (h *Handler) patchOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["orderId"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid orderId"))
	}
	data, validations := validators.PatchOrder(r)
	if validations != nil {
		throwValidationError(w, validations)
		return
	}
	order, err := h.Service.EditOrder(uint(id), data)
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(order); err != nil {
		throwInternalError(w, err)
		return
	}
}
