package usersPresentation

import (
	"encoding/json"
	"net/http"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	usersApplication "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	usersDomain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
)

func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req usersApplication.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(w, http.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return
	}

	response, err := h.loginService.Login(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err == usersDomain.ErrInvalidCredentials {
			status = http.StatusUnauthorized
		}
		commonPresentation.JsonResponseWithStatus(w, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(w, http.StatusOK, commonDomain.SuccessDataResult("user_logged_in_successfully", response))
}
