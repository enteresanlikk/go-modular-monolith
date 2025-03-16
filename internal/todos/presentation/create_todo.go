package todos_presentation

import (
	"encoding/json"
	"net/http"

	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	common_presentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todos_application "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
)

func (s *TodosHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req todos_application.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common_presentation.JsonResponseWithStatus(w, http.StatusBadRequest, common_domain.ErrorResult(err.Error()))
		return
	}

	response, err := s.createTodoService.CreateTodo(&req)
	if err != nil {
		status := http.StatusInternalServerError
		common_presentation.JsonResponseWithStatus(w, status, common_domain.ErrorResult(err.Error()))
		return
	}

	common_presentation.JsonResponseWithStatus(w, http.StatusOK, common_domain.SuccessDataResult("todo_created_successfully", response))
}
