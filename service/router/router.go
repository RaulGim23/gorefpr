package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"files/service"
)

// NewMuxRouter godoc.
func NewMuxRouter(router *mux.Router, logger service.Logger) service.Router {
	return &muxRouter{
		mux: router,
		log: logger,
	}
}

type muxRouter struct {
	mux *mux.Router
	log service.Logger
}

// Get godoc.
func (router *muxRouter) Get(path string, handler service.RouteHandler) {
	router.mux.HandleFunc(path, router.serviceHanlder(handler)).Methods(http.MethodGet)
}

// Post godoc.
func (router *muxRouter) Post(path string, handler service.RouteHandler) {
	router.mux.HandleFunc(path, router.serviceHanlder(handler)).Methods(http.MethodPost)
}

// Put godoc.
func (router *muxRouter) Put(path string, handler service.RouteHandler) {
	router.mux.HandleFunc(path, router.serviceHanlder(handler)).Methods(http.MethodPut)
}

// Delete godoc.
func (router *muxRouter) Delete(path string, handler service.RouteHandler) {
	router.mux.HandleFunc(path, router.serviceHanlder(handler)).Methods(http.MethodDelete)
}

func (router *muxRouter) serviceHanlder(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			router.log.Debugf("%s (%s): %s", r.URL.EscapedPath(), r.RemoteAddr, err)
		}
	}
}
