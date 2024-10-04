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

func (uc *TodoUseCase) Create(task *entity.Todo) (*entity.Todo, error) {

	task, err := uc.repo.Insert(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TodoUseCase) FindAll() ([]*entity.Todo, error) {

	todos, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (uc *TodoUseCase) FindById(id string) (*entity.Todo, error) {
	task, err := uc.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TodoUseCase) Update(task *entity.Todo) (*entity.Todo, error) {
	task, err := uc.repo.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TodoUseCase) Delete(id string) error {
	err := uc.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
