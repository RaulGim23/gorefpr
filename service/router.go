package service

import (
	"net/http"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request) error

type Router interface {
	Get(path string, handler RouteHandler)
	Post(path string, handler RouteHandler)
	Put(path string, handler RouteHandler)
	Delete(path string, handler RouteHandler)
}
