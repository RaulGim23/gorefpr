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

func NewFile(router service.Router, fileService service.File, logger service.Logger) {
	handler := file{
		svc: fileService,
		log: logger,
	}
	router.Get("/users", handler.List)
	router.Get("/users/{id}", handler.Get)
}

func (h *file) List(w http.ResponseWriter, r *http.Request) error {
	files, err := response.FromFilesModel(h.svc.FindAll(r.Context()))
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, "no entries")
	}
	return response.NewList(w, http.StatusOK, files)
}

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