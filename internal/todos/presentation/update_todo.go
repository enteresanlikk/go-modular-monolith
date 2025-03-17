package todosPresentation

import (
	"encoding/json"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func (s *TodosHandler) UpdateTodo(ctx *fasthttp.RequestCtx) {
	idParam := ctx.UserValue("id").(string)
	id, err := uuid.Parse(idParam)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusBadRequest, commonDomain.ErrorResult("invalid id format"))
		return
	}

	var req todosApplication.UpdateTodoRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return
	}
	req.ID = id

	response, err := s.updateTodoService.UpdateTodo(&req)
	if err != nil {
		status := fasthttp.StatusInternalServerError
		commonPresentation.JsonResponseWithStatus(ctx, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusOK, commonDomain.SuccessDataResult("todo_updated_successfully", response))
}
