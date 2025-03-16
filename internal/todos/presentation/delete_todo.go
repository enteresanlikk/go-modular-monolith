package todos_presentation

import (
	"net/http"

	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	common_presentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todos_application "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *TodosHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	req := todos_application.DeleteTodoRequest{
		ID: uuid.MustParse(mux.Vars(r)["id"]),
	}

	err := s.deleteTodoService.DeleteTodo(&req)
	if err != nil {
		status := http.StatusInternalServerError
		common_presentation.JsonResponseWithStatus(w, status, common_domain.ErrorResult(err.Error()))
		return
	}

	common_presentation.JsonResponseWithStatus(w, http.StatusOK, common_domain.SuccessResult("todo_deleted_successfully"))
}
