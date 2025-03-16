package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/enteresanlikk/go-modular-monolith/internal/common"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
)

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req application.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.JsonResponseWithStatus(w, http.StatusBadRequest, common.ErrorDataResult("Invalid request", err))
		return
	}

	response, err := h.loginService.Login(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == users_domain.ErrInvalidCredentials {
			status = http.StatusUnauthorized
		}
		common.JsonResponseWithStatus(w, status, common.ErrorDataResult("Failed to login", err))
		return
	}

	common.JsonResponseWithStatus(w, http.StatusOK, common.SuccessDataResult("User logged in successfully", response))
}
