package users_presentation

import (
	"encoding/json"
	"net/http"

	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	common_presentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	users_application "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
)

func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req users_application.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common_presentation.JsonResponseWithStatus(w, http.StatusBadRequest, common_domain.ErrorDataResult("Invalid request", err))
		return
	}

	response, err := h.loginService.Login(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == users_domain.ErrInvalidCredentials {
			status = http.StatusUnauthorized
		}
		common_presentation.JsonResponseWithStatus(w, status, common_domain.ErrorDataResult("Failed to login", err))
		return
	}

	common_presentation.JsonResponseWithStatus(w, http.StatusOK, common_domain.SuccessDataResult("User logged in successfully", response))
}
