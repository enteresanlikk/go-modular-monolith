package presentation

import (
	"net/http"

	"github.com/enteresanlikk/go-modular-monolith/internal/common"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/application/users/login_user"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/application/users/register_user"
	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain/users"
	users_infrastructure "github.com/enteresanlikk/go-modular-monolith/internal/users/infrastructure/users"
	"github.com/gin-gonic/gin"
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

func (h *AuthHandler) Register(c *gin.Context) {
	var req register_user.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorDataResult("Invalid request", err.Error()))
		return
	}

	response, err := h.registerService.Register(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == users_domain.ErrPasswordMismatch {
			status = http.StatusBadRequest
		}
		c.JSON(status, common.ErrorDataResult("Failed to register user", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, common.SuccessDataResult("User created successfully", response))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req login_user.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorDataResult("Invalid request", err.Error()))
		return
	}

	response, err := h.loginService.Login(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == users_domain.ErrInvalidCredentials {
			status = http.StatusUnauthorized
		}
		c.JSON(status, common.ErrorDataResult("Failed to login", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.SuccessDataResult("User logged in successfully", response))
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := users_infrastructure.NewUserRepository(db)
	registerService := register_user.NewUserService(userRepo)
	loginService := login_user.NewUserService(userRepo)
	handler := NewAuthHandler(registerService, loginService)

	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
}
