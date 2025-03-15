package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/enteresanlikk/go-modular-monolith/internal/common"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/application/users/login_user"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/application/users/register_user"
	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain/users"
	users_infrastructure "github.com/enteresanlikk/go-modular-monolith/internal/users/infrastructure/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AuthHandler struct {
	registerService register_user.UserService
	loginService    login_user.UserService
}

func NewAuthHandler(registerService *register_user.UserService, loginService *login_user.UserService) *AuthHandler {
	return &AuthHandler{
		registerService: *registerService,
		loginService:    *loginService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req register_user.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response, err := h.registerService.Register(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == users_domain.ErrPasswordMismatch {
			status = http.StatusBadRequest
		}
		http.Error(w, "Failed to register user", status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(common.SuccessDataResult("User created successfully", response))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req login_user.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response, err := h.loginService.Login(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == users_domain.ErrInvalidCredentials {
			status = http.StatusUnauthorized
		}
		http.Error(w, "Failed to login", status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(common.SuccessDataResult("User logged in successfully", response))
}

func RegisterRoutes(mux *mux.Router, db *gorm.DB) {
	userRepo := users_infrastructure.NewUserRepository(db)
	registerService := register_user.NewUserService(userRepo)
	loginService := login_user.NewUserService(userRepo)
	handler := NewAuthHandler(registerService, loginService)

	authRouter := mux.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
}
