package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"effectiveMobileTest/internal/dto"
)

type Handler struct {
	service service
	router  *mux.Router
}

func New(service service) *Handler {
	h := Handler{service: service, router: mux.NewRouter()}

	h.router.Path("/persons").Methods(http.MethodGet).HandlerFunc(h.GetPersons)
	h.router.Path("/persons").Methods(http.MethodPost).HandlerFunc(h.CreatePerson)
	h.router.Path("/persons/{userID:[0-9]+}").Methods(http.MethodDelete).HandlerFunc(h.DeletePerson)
	h.router.Path("/persons/{userID:[0-9]+}").Methods(http.MethodPut).HandlerFunc(h.UpdatePerson)

	return &h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

type errorResponse struct {
	Error string `json:"error"`
}

func responseHandler(response interface{}, err error, w http.ResponseWriter) {
	if err == nil {
		json.NewEncoder(w).Encode(response)

		return
	}

	statusCode := http.StatusInternalServerError

	switch {
	case errors.Is(err, dto.ErrBadRequest):
		statusCode = http.StatusBadRequest
	case errors.Is(err, dto.ErrNotFound):
		statusCode = http.StatusNotFound
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse{
		Error: err.Error(),
	})
}
