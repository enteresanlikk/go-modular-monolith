package users

import (
	"net/http"

	users_application "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	users_infrastructure "github.com/enteresanlikk/go-modular-monolith/internal/users/infrastructure"
	users_presentation "github.com/enteresanlikk/go-modular-monolith/internal/users/presentation"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Register(mux *mux.Router, db *gorm.DB) {
	userRepo := users_infrastructure.NewUserRepository(db)
	registerService := users_application.NewUserService(userRepo)
	loginService := users_application.NewUserService(userRepo)
	handler := users_presentation.NewUsersHandler(registerService, loginService)

	authRouter := mux.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
}
