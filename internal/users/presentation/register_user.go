package usersPresentation

import (
	"github.com/gofiber/fiber/v2"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	usersApplication "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
)

func (h *UsersHandler) Register(ctx *fiber.Ctx) error {
	var req usersApplication.RegisterUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	response, err := h.registerService.Register(&req)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusCreated, commonDomain.SuccessDataResult("user_registered_successfully", response))
	return nil
}
