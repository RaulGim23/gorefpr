package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"files/service"
	"files/transport/request"
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
	router.Post("/files", handler.Create)
	router.Put("/files/{id}", handler.Update)
	router.Delete("/files/{id}", handler.Delete)
}

// List godoc.
func (h *file) List(w http.ResponseWriter, r *http.Request) error {
	order := r.URL.Query().Get("order")
	orderBy := ToSnakeCase(r.URL.Query().Get("orderBy"))
	var orderBys []string
	if order != "" && orderBy != "" {
		orderBys = []string{orderBy, order}
	}
	page, err := uintFromQuery(r, "page", 0)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "invalid page value")
	}

	limit, err := uintFromQuery(r, "limit", 10)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "invalid limit value")
	}

	files, total, err := response.FromFilesModel(h.svc.FindAll(r.Context(), orderBys, page, limit))
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, "no entries")
	}

	return response.NewList(w, http.StatusOK, page, limit, total, files)
}

// Get godoc.
func (h *file) Get(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	fileID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "invalid file id")
	}

	device, err := h.svc.Find(r.Context(), fileID)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, "internal error")
	}

	return response.New(w, http.StatusOK, response.FromFile(device))
}

// Create godoc.
func (h *file) Create(w http.ResponseWriter, r *http.Request) error {
	file, err := request.FileFromPayload(r.Body)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "malformed request")
	}
	date, err := time.Parse("YYYY-MM-DD", file.Date)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "The date should be of type YYYY-MM-DD")
	}
	fmt.Println(date)
	err = h.svc.Store(r.Context(), file)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, "internal error")
	}
	return response.New(w, http.StatusOK, response.FromFile(file))
}

// Update godoc.
func (h *file) Update(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	fileID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "invalid file id")
	}
	file, err := request.FileFromPayload(r.Body)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "malformed request")
	}
	file.ID = fileID
	err = h.svc.Update(r.Context(), file)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, "internal error")
	}
	updatedFile, err := h.svc.Find(r.Context(), file.ID)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, "internal error")
	}
	return response.New(w, http.StatusOK, response.FromFile(updatedFile))
}

// Delete godoc.
func (h *file) Delete(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	fileID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return response.NewError(w, http.StatusNotFound, "File not found.")
	}

	err = h.svc.Delete(r.Context(), fileID)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, "internal error")
	}
	deletedFile, err := h.svc.Find(r.Context(), fileID)
	return response.JSON(w, http.StatusOK, deletedFile)
}
