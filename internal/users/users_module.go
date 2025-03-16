package users

import (
	"net/http"
	"os"
	"time"

	usersApplication "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	usersDomain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	usersInfrastructure "github.com/enteresanlikk/go-modular-monolith/internal/users/infrastructure"
	usersPresentation "github.com/enteresanlikk/go-modular-monolith/internal/users/presentation"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Register(mux *mux.Router, db *gorm.DB) {
	userRepo := usersInfrastructure.NewUserRepository(db)

	tokenConfig := usersDomain.TokenConfig{
		AccessTokenDuration:  time.Hour * 1,
		RefreshTokenDuration: time.Hour * 24 * 7,
		AccessTokenSecret:    []byte(os.Getenv("JWT_SECRET")),
		RefreshTokenSecret:   []byte(os.Getenv("JWT_REFRESH_SECRET")),
	}
	tokenService := usersInfrastructure.NewTokenService(tokenConfig)

	registerService := usersApplication.NewUserService(userRepo, tokenService)
	loginService := usersApplication.NewUserService(userRepo, tokenService)

	handler := usersPresentation.NewUsersHandler(registerService, loginService)

	authRouter := mux.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
}
