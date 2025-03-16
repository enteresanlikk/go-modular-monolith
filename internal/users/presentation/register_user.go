package users_presentation

import (
	"encoding/json"
	"net/http"

	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	common_presentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	users_application "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
)

func (h *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req users_application.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common_presentation.JsonResponseWithStatus(w, http.StatusBadRequest, common_domain.ErrorDataResult("Invalid request", err))
		return
	}

	response, err := h.registerService.Register(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == users_domain.ErrPasswordMismatch {
			status = http.StatusBadRequest
		}
		common_presentation.JsonResponseWithStatus(w, status, common_domain.ErrorDataResult("Failed to register user", err))
		return
	}

	common_presentation.JsonResponseWithStatus(w, http.StatusCreated, common_domain.SuccessDataResult("User created successfully", response))
}
