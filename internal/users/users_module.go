package users

import (
	"net/http"
	"os"
	"time"

	users_application "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	users_infrastructure "github.com/enteresanlikk/go-modular-monolith/internal/users/infrastructure"
	users_presentation "github.com/enteresanlikk/go-modular-monolith/internal/users/presentation"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Register(mux *mux.Router, db *gorm.DB) {
	// Initialize repositories
	userRepo := users_infrastructure.NewUserRepository(db)

	// Initialize token service
	tokenConfig := users_domain.TokenConfig{
		AccessTokenDuration:  time.Hour * 1,      // 1 hour
		RefreshTokenDuration: time.Hour * 24 * 7, // 7 days
		AccessTokenSecret:    []byte(os.Getenv("JWT_SECRET")),
		RefreshTokenSecret:   []byte(os.Getenv("JWT_REFRESH_SECRET")),
	}
	tokenService := users_infrastructure.NewTokenService(tokenConfig)

	// Initialize user services
	registerService := users_application.NewUserService(userRepo, tokenService)
	loginService := users_application.NewUserService(userRepo, tokenService)

	// Initialize handlers
	handler := users_presentation.NewUsersHandler(registerService, loginService)

	// Register routes
	authRouter := mux.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
}
