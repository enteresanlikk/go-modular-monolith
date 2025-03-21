package todosPresentation

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *TodosHandler) GetTodoById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusBadRequest, commonDomain.ErrorResult("invalid id format"))
		return nil
	}

	req := todosApplication.GetTodoByIdRequest{
		ID: id,
	}

	response, err := s.getTodoByIdService.GetTodoById(&req)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusInternalServerError, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusOK, commonDomain.SuccessDataResult("todo_fetched_successfully", response))
	return nil
}
