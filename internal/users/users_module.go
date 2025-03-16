package users

import (
	"net/http"

	"github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	users_infrastructure "github.com/enteresanlikk/go-modular-monolith/internal/users/infrastructure"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/presentation"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(mux *mux.Router, db *gorm.DB) {
	userRepo := users_infrastructure.NewUserRepository(db)
	registerService := application.NewUserService(userRepo)
	loginService := application.NewUserService(userRepo)
	handler := presentation.NewAuthHandler(registerService, loginService)

	authRouter := mux.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
}
