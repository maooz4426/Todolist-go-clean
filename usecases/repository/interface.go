package repository

import (
	"context"
	"github.com/maooz4426/Todolist/domain/entity"
)

// usecase層で使うrepository層のメソッドを定義
// これで依存先を抽象化
// ダックタイピングしてる
type ITodoRepository interface {
	Insert(ctx context.Context, task *entity.Todo) (*entity.Todo, error)
	FindAll(ctx context.Context) ([]*entity.Todo, error)
	FindById(ctx context.Context, id string) (*entity.Todo, error)
	Update(ctx context.Context, task *entity.Todo) (*entity.Todo, error)
	Delete(ctx context.Context, id string) error
}
