package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"files/service"
	"files/transport/response"
)

type file struct {
	svc service.File
	log service.Logger
}

// NewFile godoc.
func NewFile(router service.Router, fileService service.File, logger service.Logger) {
	handler := file{
		svc: fileService,
		log: logger,
	}
	router.Get("/files", handler.List)
	router.Get("/files/{id}", handler.Get)
}

// List godoc.
func (h *file) List(w http.ResponseWriter, r *http.Request) error {
	urlOrder := r.URL.Query().Get("order")
	urlOrderBy := r.URL.Query().Get("orderBy")
	page, err := uintFromQuery(r, "page", 0)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "invalid page value")
	}

	limit, err := uintFromQuery(r, "limit", 10)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "invalid limit value")
	}

	files, total, err := response.FromFilesModel(h.svc.FindAll(r.Context(), urlOrder, urlOrderBy, page, limit))
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, "no entries")
	}

	return response.NewList(w, http.StatusOK, page, limit, total, files)
}

// Get godoc.
func (h *file) Get(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "invalid user id")
	}

	device, err := h.svc.Find(r.Context(), userID)
	if err != nil {
		h.log.Debugf("user.Get %s: %s", r.RemoteAddr, err)

		return response.NewError(w, http.StatusInternalServerError, "internal error")
	}

	return response.New(w, http.StatusOK, response.FromFile(device))
}
