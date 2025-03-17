package todosPresentation

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func (s *TodosHandler) GetTodoById(ctx *fasthttp.RequestCtx) {
	idParam := ctx.UserValue("id").(string)
	id, err := uuid.Parse(idParam)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusBadRequest, commonDomain.ErrorResult("invalid id format"))
		return
	}

	req := todosApplication.GetTodoByIdRequest{
		ID: id,
	}

	response, err := s.getTodoByIdService.GetTodoById(&req)
	if err != nil {
		status := fasthttp.StatusInternalServerError
		commonPresentation.JsonResponseWithStatus(ctx, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusOK, commonDomain.SuccessDataResult("todo_fetched_successfully", response))
}
