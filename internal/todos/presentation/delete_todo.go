package todosPresentation

import (
	"net/http"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *TodosHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	req := todosApplication.DeleteTodoRequest{
		ID: uuid.MustParse(mux.Vars(r)["id"]),
	}

	err := s.deleteTodoService.DeleteTodo(&req)
	if err != nil {
		status := http.StatusInternalServerError
		commonPresentation.JsonResponseWithStatus(w, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(w, http.StatusOK, commonDomain.SuccessResult("todo_deleted_successfully"))
}
