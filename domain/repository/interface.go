package repository

import (
	"context"
	"github.com/maooz4426/Todolist/domain"
)

// usecase層で使うrepository層のメソッドを定義
// これで依存先を抽象化
type TodoRepository interface {
	InsertTodo(ctx context.Context, task *domain.Todo) (domain.Todo, error)
}
