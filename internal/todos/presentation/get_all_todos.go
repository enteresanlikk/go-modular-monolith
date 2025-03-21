package todosPresentation

import (
	"github.com/gofiber/fiber/v2"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
)

func (s *TodosHandler) GetAllTodos(ctx *fiber.Ctx) error {
	var req todosApplication.GetAllTodosRequest
	if err := ctx.BodyParser(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	response, err := s.getAllTodosService.GetAllTodos(&req)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusInternalServerError, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	commonPresentation.JsonResponseWithStatus(ctx, fiber.StatusOK, commonDomain.SuccessDataResult("todos_fetched_successfully", response))
	return nil
}
