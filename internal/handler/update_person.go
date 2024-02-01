package handler

import (
	"effectiveMobileTest/internal/dto"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h *Handler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slog.Debug("Start UpdatePerson")
	defer slog.Debug("End UpdatePerson")

	request := dto.UpdatePersonRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "error parsing json", http.StatusBadRequest)

		return
	}

	id, err := PathVarInt(r, "userID")
	if err != nil {
		http.Error(w, "error parse userID", http.StatusBadRequest)

		return
	}

	response, err := h.service.UpdatePerson(&request, id)
	responseHandler(response, err, w)
}
