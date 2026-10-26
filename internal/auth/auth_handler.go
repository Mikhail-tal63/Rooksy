package auth

import (
	"net/http"

	errorhandler "github.com/Mikhail-Tal63/Rooksy/util/errorHandler"
	"github.com/Mikhail-Tal63/Rooksy/util/jso"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	Service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Router(mux *mux.Router) {
	mux.HandleFunc("/register", h.Register).Methods("POST")
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload *RegisterRequist

	if err := jso.ParseJSON(r, payload); err != nil {
		errorhandler.Handle(w, err)
		return
	}

	res, err := h.Service.Register(r.Context(), *payload)
	if err != nil {
		errorhandler.Handle(w, err)
		return
	}
	if err := jso.WriteJSON(w, http.StatusCreated, map[string]any{
		"message":  "user registerd seccssfully",
		"responce": res,
	}); err != nil {
		errorhandler.Handle(w, err)
		return
	}
}
