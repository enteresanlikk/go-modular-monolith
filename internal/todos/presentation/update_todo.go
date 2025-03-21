package todosPresentation

import (
	"github.com/gofiber/fiber/v2"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/google/uuid"
)

func (s *TodosHandler) UpdateTodo(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusBadRequest, commonDomain.ErrorResult("invalid id format"))
		return nil
	}

	var req todosApplication.UpdateTodoRequest
	if err := ctx.BodyParser(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return nil
	}
	req.ID = id

	response, err := s.updateTodoService.UpdateTodo(&req)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusInternalServerError, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusOK, commonDomain.SuccessDataResult("todo_updated_successfully", response))
	return nil
}
