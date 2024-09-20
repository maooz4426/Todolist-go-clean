package repository

import (
	"github.com/maooz4426/Todolist/domain/entity"
)

// usecase層で使うrepository層のメソッドを定義
// これで依存先を抽象化
type TodoRepositoryer interface {
	InsertTodo(task *entity.Todo) (*entity.Todo, error)
}
