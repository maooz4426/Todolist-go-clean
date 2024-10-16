package port

import (
	"context"
	"github.com/maooz4426/Todolist/domain/entity"
)

type ITodoUseCase interface {
	Create(ctx context.Context, task *entity.Todo) (*entity.Todo, error)
	FindAll(ctx context.Context) ([]*entity.Todo, error)
	FindById(ctx context.Context, id string) (*entity.Todo, error)
	Update(ctx context.Context, task *entity.Todo) (*entity.Todo, error)
	Delete(ctx context.Context, id string) (*entity.Todo, error)
}
