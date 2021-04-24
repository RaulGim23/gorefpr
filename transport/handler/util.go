package handler

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase godoc.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func uintFromQuery(r *http.Request, param string, defaultValue uint64) (uint64, error) {
	// Get limit param
	limitStr := r.URL.Query().Get(param)
	if limitStr != "" {
		return strconv.ParseUint(limitStr, 10, 64)
	}

	return defaultValue, nil
}
