package http

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func throwValidationError(w http.ResponseWriter, validations url.Values) {
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(validations)
}

func throwInternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
