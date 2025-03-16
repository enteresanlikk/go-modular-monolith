package todosPresentation

import (
	"encoding/json"
	"net/http"

	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *TodosHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	req := todosApplication.UpdateTodoRequest{
		ID: uuid.MustParse(mux.Vars(r)["id"]),
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(w, http.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return
	}

	response, err := s.updateTodoService.UpdateTodo(&req)
	if err != nil {
		status := http.StatusInternalServerError
		commonPresentation.JsonResponseWithStatus(w, status, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(w, http.StatusOK, commonDomain.SuccessDataResult("todo_updated_successfully", response))
}
