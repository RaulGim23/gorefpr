package handler

import (
	"net/http"
	"strconv"
)

func uintFromQuery(r *http.Request, param string, defaultValue uint64) (uint64, error) {
	// Get limit param
	limitStr := r.URL.Query().Get(param)
	if limitStr != "" {
		return strconv.ParseUint(limitStr, 10, 64)
	}
	return defaultValue, nil
}
