package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"effectiveMobileTest/internal/dto"
)

func (h *Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slog.Debug("Start CreatePerson")
	defer slog.Debug("End CreatePerson")

	request := dto.CreatePersonRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "error parsing json", http.StatusBadRequest)

		return
	}

	response, err := h.service.CreatePerson(&request)
	responseHandler(response, err, w)
}
