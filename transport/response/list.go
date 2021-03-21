package response

import (
	"encoding/json"
	"net/http"
)

// NewList godoc
func NewList(w http.ResponseWriter, status int, files []File) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(files)
}

func New(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}