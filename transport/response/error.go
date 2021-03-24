package response

import (
	"encoding/json"
	"net/http"
)

// Error error response payload.
type Error struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// NewError create a new HTTP error and send it trough http.ResponseWriter.
func NewError(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(Error{
		Code:    status,
		Message: message,
	})
}
