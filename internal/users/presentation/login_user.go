package usersPresentation

import (
	"github.com/goccy/go-json"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	usersApplication "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	"github.com/valyala/fasthttp"
)

func (h *UsersHandler) Login(ctx *fasthttp.RequestCtx) {
	var req usersApplication.LoginUserRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return
	}

	response, err := h.loginService.Login(&req)
	if err != nil {
		status := fasthttp.StatusBadRequest
		commonPresentation.JsonResponseWithStatus(ctx, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusOK, commonDomain.SuccessDataResult("user_logged_in_successfully", response))
}
