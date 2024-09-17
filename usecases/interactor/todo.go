package interactor

import (
	"github.com/maooz4426/Todolist/domain/entity"
	"github.com/maooz4426/Todolist/usecases/repository"
)

type TodoUseCase struct {
	repo repository.TodoRepositoryer
}

func NewTodoUseCase(repo repository.TodoRepositoryer) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

func (uc *TodoUseCase) Create(task *entity.Todo) error {

	uc.repo.InsertTodo(task)
	return nil
}
