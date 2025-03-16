package todosPresentation

import (
	"encoding/json"
	"net/http"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
)

func (s *TodosHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	var req todosApplication.GetAllTodosRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(w, http.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return
	}

	response, err := s.getAllTodosService.GetAllTodos(&req)
	if err != nil {
		status := http.StatusInternalServerError
		commonPresentation.JsonResponseWithStatus(w, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(w, http.StatusOK, commonDomain.SuccessDataResult("todos_fetched_successfully", response))
}
