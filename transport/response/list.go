package response

import (
	"encoding/json"
	"net/http"
)

// List paginated list
type List struct {
	Page    uint64      `json:"page"`
	Limit   uint64      `json:"limit"`
	Total   uint64      `json:"total"`
	Results interface{} `json:"results"`
}

// NewList godoc.
func NewList(w http.ResponseWriter, status int, page, limit, total uint64, results []File) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(List{
		Page:    page,
		Limit:   limit,
		Total:   total,
		Results: results,
	})
}

// New return new Json response.
func New(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
