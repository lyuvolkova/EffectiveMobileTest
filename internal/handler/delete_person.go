package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"effectiveMobileTest/internal/dto"
)

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slog.Debug("Start DeletePerson")
	defer slog.Debug("End DeletePerson")
	request := dto.DeletePersonRequest{}

	var err error

	request.ID, err = PathVarInt(r, "userID")
	if err != nil {
		http.Error(w, "error parse userID", http.StatusBadRequest)

		return
	}
	response, err := h.service.DeletePerson(&request)
	responseHandler(response, err, w)
}

func PathVarInt(r *http.Request, name string) (int, error) {
	value := mux.Vars(r)[name]

	return strconv.Atoi(value)
}
