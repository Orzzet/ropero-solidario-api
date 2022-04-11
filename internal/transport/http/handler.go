package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/orzzet/ropero-solidario-api/src/services"
	"net/http"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Secret string
	Router *mux.Router
	*services.Service
}

type Response struct {
	Message string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(db *gorm.DB, secret string) *Handler {
	return &Handler{
		Secret:  secret,
		Service: services.NewService(db),
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.Use(cors)
	h.Router.HandleFunc("/auth", h.createToken).Methods("POST", "OPTIONS")

	// Orders
	h.Router.HandleFunc("/orders", h.getOrders).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/orders", h.createOrder).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/orders/{orderId}", h.getOrder).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/orders/{orderId}", h.patchOrder).Methods("PATCH", "OPTIONS")

	// Items
	h.Router.HandleFunc("/items", h.getItems).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/items", h.createItem).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/items/{itemId}", h.editItem).Methods("PUT", "OPTIONS")
	h.Router.HandleFunc("/items/{itemId}", h.deleteItem).Methods("DELETE", "OPTIONS")

	// Categories
	h.Router.HandleFunc("/categories", h.auth(h.getCategories)).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/categories", h.getCategories).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/categories/bulk", h.createCategories).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/categories/{categoryId}", h.deleteCategory).Methods("DELETE", "OPTIONS")
	h.Router.HandleFunc("/categories/{categoryId}", h.editCategory).Methods("PUT", "OPTIONS")

	// Users
	h.Router.HandleFunc("/users", h.createUser).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/users", h.getUsers).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/users/{userId}", h.getUser).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/users/{userId}", h.deleteUser).Methods("DELETE", "OPTIONS")
	h.Router.HandleFunc("/users/{userId}/approve", h.approveUser).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/users/{userId}/resetPassword", h.resetUserPassword).Methods("POST", "OPTIONS")

	// Health check
	h.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(Response{Message: "Online"}); err != nil {
			panic(err)
		}
	})
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}
