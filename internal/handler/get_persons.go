package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"effectiveMobileTest/internal/dto"
)

func (h *Handler) GetPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slog.Debug("Start GetPersons")
	defer slog.Debug("End GetPersons")

	request := dto.GetPersonsRequest{
		Name:       r.URL.Query().Get("name"),
		Surname:    r.URL.Query().Get("surname"),
		Patronymic: r.URL.Query().Get("patronymic"),
		Page:       1,
	}

	page := r.URL.Query().Get("page")
	var err error
	if page != "" {
		request.Page, err = strconv.Atoi(page)
		if err != nil || request.Page <= 0 {
			responseHandler(nil, fmt.Errorf("page: %w", dto.ErrBadRequest), w)
			return
		}
	}

	response, err := h.service.GetPersons(&request)
	responseHandler(response, err, w)
}
