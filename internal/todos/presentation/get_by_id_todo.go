package todosPresentation

import (
	"net/http"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *TodosHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	req := todosApplication.GetTodoByIdRequest{
		ID: uuid.MustParse(mux.Vars(r)["id"]),
	}

	response, err := s.getTodoByIdService.GetTodoById(&req)
	if err != nil {
		status := http.StatusInternalServerError
		commonPresentation.JsonResponseWithStatus(w, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(w, http.StatusOK, commonDomain.SuccessDataResult("todo_fetched_successfully", response))
}
