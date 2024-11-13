package interactor

import (
	"context"
	"errors"
	"fmt"
	"github.com/maooz4426/Todolist/domain/entity"
	repository "github.com/maooz4426/Todolist/domain/repository"
	"strconv"
)

type TodoUseCase struct {
	repo repository.ITodoRepository
	txm  repository.ITransactionManager
}

func NewTodoUseCase(repo repository.ITodoRepository, txm repository.ITransactionManager) *TodoUseCase {
	return &TodoUseCase{repo: repo, txm: txm}
}

func (uc *TodoUseCase) Create(ctx context.Context, taskReq *entity.Todo) (*entity.Todo, error) {
	var task *entity.Todo
	var err error

	err = uc.txm.RunInTx(ctx, func(ctx context.Context) error {
		task, err = uc.repo.Insert(ctx, taskReq)
		if err != nil {
			return err
		}

		id := strconv.Itoa(int(task.ID))

		var confirm *entity.Todo

		confirm, err = uc.repo.FindById(ctx, id)
		if err != nil {
			return err
		}

		if confirm == nil {
			return errors.New("todo not found")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TodoUseCase) FindAll(ctx context.Context) ([]*entity.Todo, error) {

	todos, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (uc *TodoUseCase) FindById(ctx context.Context, id string) (*entity.Todo, error) {
	task, err := uc.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TodoUseCase) Update(ctx context.Context, task *entity.Todo) (*entity.Todo, error) {

	err := uc.txm.RunInTx(ctx, func(ctx context.Context) error {
		exsist, err := uc.repo.FindById(ctx, strconv.Itoa(int(task.ID)))
		if err != nil {
			return err
		}
		if exsist == nil {
			return errors.New("todo not found")
		}

		task, err = uc.repo.Update(ctx, task)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TodoUseCase) Delete(ctx context.Context, id string) (*entity.Todo, error) {

	var todo *entity.Todo
	var err error

	err = uc.txm.RunInTx(ctx, func(ctx context.Context) error {
		todo, err = uc.repo.FindById(ctx, id)
		fmt.Println(todo)
		if err != nil {
			return err
		}

		if todo == nil {
			return errors.New("todo not found") //ここなに?
		}

		err = uc.repo.Delete(ctx, id)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return todo, nil
}
