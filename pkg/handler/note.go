package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorvelo/notes/internal/models"
)

const (
	NotFoundID     = "Not found ID"
	InvalidIdParam = "invalid id param"
	InvalidRequest = "Invalid resquest payload"
	NoteDelete     = "Note was deleted!"
)

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	var noteInput models.Note

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&noteInput); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	id, ok := JWTClaimsFromContext(r.Context())
	if !ok {
		ResponseError(w, http.StatusInternalServerError, NotFoundID)
	}

	body, err := CheckNote(noteInput.Description)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	if err := h.service.Notes.Add(id, noteInput); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(body) == 0 {
		ResponseJSON(w, http.StatusCreated, noteInput.Description)
	} else {
		ResponseJSON(w, http.StatusCreated, body)
	}
}

func (h *Handler) getAllNotes(w http.ResponseWriter, r *http.Request) {
	id, ok := JWTClaimsFromContext(r.Context())
	if !ok {
		ResponseError(w, http.StatusInternalServerError, NotFoundID)
	}
	notes, err := h.service.Notes.ListAll(id)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJSON(w, http.StatusOK, notes)
}

func (h *Handler) deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseError(w, http.StatusBadRequest, InvalidIdParam)
		return
	}

	if err := h.service.Notes.Delete(id); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, NoteDelete)
}

func (h *Handler) upgradeNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseError(w, http.StatusBadRequest, InvalidIdParam)
		return
	}

	var input models.Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, InvalidRequest)
		return
	}
	defer r.Body.Close()
	input.Id = id

	if err = h.service.Notes.Update(id, &input); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, input.Description)
}
