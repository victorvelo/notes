package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorvelo/notes/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home paige!")
	})

	auth := router.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
	auth.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(h.AuthenticationMW)
	protected.HandleFunc("/notes", h.createNote).Methods(http.MethodPost)
	protected.HandleFunc("/notes/{id:[0-9]+}", h.deleteNote).Methods(http.MethodDelete)
	protected.HandleFunc("/notes/{id:[0-9]+}", h.upgradeNote).Methods(http.MethodPut)
	protected.HandleFunc("/notes/all", h.getAllNotes).Methods(http.MethodGet)

	return router
}
