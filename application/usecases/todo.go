package usecases

import (
	"context"
	"github.com/maooz4426/Todolist/domain"
	"github.com/maooz4426/Todolist/domain/repository"
)

type TodoUseCase struct {
	repo repository.TodoRepository
}

func NewTodoUseCase(repo repository.TodoRepository) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

func (uc *TodoUseCase) Add(ctx context.Context, task *domain.Todo) error {

	uc.repo.InsertTodo(ctx, task)
	return nil
}
