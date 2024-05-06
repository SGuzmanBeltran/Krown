package user

import (
	"championForge/types"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	userStore *Store
}

func NewHandler(userStore *Store) *Handler {
	return &Handler{userStore}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Route("/api/user", func(r chi.Router) {
		r.Post("/register", h.handleRegister)
		r.Post("/login", h.handleLogin)
		r.Get("/", h.getUsers)
	})
}

func (h *Handler)handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")
}

func (h *Handler)handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	err := json.NewDecoder(r.Body).Decode(payload)
	// todo: check if user exists

	// todo: create it
	fmt.Println("Login")
}

func (h *Handler)getUsers(w http.ResponseWriter, r *http.Request) {
	// todo: get json payload
	// todo: check if user exists
	// todo: create it
	fmt.Println("users")
}