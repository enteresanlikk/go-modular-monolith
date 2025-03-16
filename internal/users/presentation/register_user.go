package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/enteresanlikk/go-modular-monolith/internal/common"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
)

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req application.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.JsonResponseWithStatus(w, http.StatusBadRequest, common.ErrorDataResult("Invalid request", err))
		return
	}

	response, err := h.registerService.Register(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == users_domain.ErrPasswordMismatch {
			status = http.StatusBadRequest
		}
		common.JsonResponseWithStatus(w, status, common.ErrorDataResult("Failed to register user", err))
		return
	}

	common.JsonResponseWithStatus(w, http.StatusCreated, common.SuccessDataResult("User created successfully", response))
}
