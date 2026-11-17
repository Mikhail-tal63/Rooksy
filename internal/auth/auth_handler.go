package auth

import (
	"log"
	"net/http"

	"github.com/Mikhail-Tal63/Rooksy/internal/middleware"
	errorhandler "github.com/Mikhail-Tal63/Rooksy/util/errorHandler"
	"github.com/Mikhail-Tal63/Rooksy/util/jso"
	"github.com/google/uuid"
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

func (h *AuthHandler) Router(mux *mux.Router, protected *mux.Router) {
	mux.HandleFunc("/register", h.Register).Methods("POST")
	mux.HandleFunc("/login", h.Login).Methods("POST")
	protected.HandleFunc("/me", h.GetME).Methods("GET")
	mux.HandleFunc("/{id:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}", h.GetUserData).Methods("GET")
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload *RegisterRequist

	if err := jso.ParseJSON(r, &payload); err != nil {
		errorhandler.Handle(w, err)
		return
	}

	res, err := h.Service.Register(r.Context(), *payload)
	if err != nil {
		errorhandler.Handle(w, err)
		log.Printf("register failed: %v", err)
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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload *LoginRequest

	if err := jso.ParseJSON(r, &payload); err != nil {
		errorhandler.Handle(w, err)
		return
	}

	authRes, err := h.Service.Login(r.Context(), *payload)
	if err != nil {
		errorhandler.Handle(w, err)
		return
	}
	if err := jso.WriteJSON(w, http.StatusOK, map[string]any{
		"message":  "welcome back",
		"responce": authRes,
	}); err != nil {
		errorhandler.Handle(w, err)
		return
	}
}

func (h *AuthHandler) GetME(w http.ResponseWriter, r *http.Request) {
	userid, err := middleware.GetUserID(r.Context())
	if err != nil {
		errorhandler.Handle(w, err)
		return
	}
	user, err := h.Service.GetUserData(r.Context(), userid)
	if err != nil {
		errorhandler.Handle(w, err)
		return
	}
	jso.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "data sent succsessfully",
		"data":    user,
	})
}
func (h *AuthHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	userid, err := uuid.Parse(id)
	if err != nil {
		errorhandler.Handle(w, ErrInvalidUserID)
		return
	}
	user, err := h.Service.GetUserData(r.Context(), userid)
	if err != nil {
		errorhandler.Handle(w, err)
		return
	}
	jso.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "data sent succsessfully",
		"data":    user,
	})
}
