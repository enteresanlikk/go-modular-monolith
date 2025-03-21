package usersPresentation

import (
	"github.com/gofiber/fiber/v2"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	usersApplication "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
)

func (h *UsersHandler) Login(ctx *fiber.Ctx) error {
	var req usersApplication.LoginUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	response, err := h.loginService.Login(&req)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusOK, commonDomain.SuccessDataResult("user_logged_in_successfully", response))
	return nil
}
