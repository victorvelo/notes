package handler

import (
	"encoding/json"
	"net/http"

	"github.com/victorvelo/notes/internal/models"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if userInput.Login == "" || userInput.Password == "" {
		ResponseError(w, http.StatusBadRequest, InvalidRequest)
		return
	}

	if err := h.service.Authorization.Add(userInput); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJSON(w, http.StatusCreated, userInput)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var userSign models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userSign); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	token, err := h.service.Authorization.CreateToken(userSign.Login, userSign.Password)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJSON(w, http.StatusCreated, token)
}
