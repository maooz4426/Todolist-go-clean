package interactor

import (
	"github.com/maooz4426/Todolist/domain/dto"
	"github.com/maooz4426/Todolist/domain/entity"
	"github.com/maooz4426/Todolist/usecases/repository"
)

type TodoUseCase struct {
	repo repository.TodoRepositoryer
}

func NewTodoUseCase(repo repository.TodoRepositoryer) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

func (uc *TodoUseCase) Create(task *entity.Todo) (*dto.CreateResponse, error) {

	task, err := uc.repo.InsertTodo(task)
	if err != nil {
		return nil, err
	}

	res := &dto.CreateResponse{
		ID:       task.ID,
		Task:     task.Task,
		Deadline: task.Deadline,
		Done:     task.Done,
	}
	
	return res, nil
}
