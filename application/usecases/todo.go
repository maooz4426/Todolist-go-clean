package usecases

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain"
	"github.com/maooz4426/Todolist/domain/repository"
)

type TodoUseCase struct {
	repo repository.TodoRepositoryer
}

func NewTodoUseCase(repo repository.TodoRepositoryer) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

func (uc *TodoUseCase) Create(ctx echo.Context, task *domain.Todo) error {

	uc.repo.InsertTodo(ctx, task)
	return nil
}
