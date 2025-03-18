package todosPresentation

import (
	"github.com/goccy/go-json"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/valyala/fasthttp"
)

func (s *TodosHandler) GetAllTodos(ctx *fasthttp.RequestCtx) {
	var req todosApplication.GetAllTodosRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return
	}

	response, err := s.getAllTodosService.GetAllTodos(&req)
	if err != nil {
		status := fasthttp.StatusInternalServerError
		commonPresentation.JsonResponseWithStatus(ctx, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(ctx, fasthttp.StatusOK, commonDomain.SuccessDataResult("todos_fetched_successfully", response))
}
