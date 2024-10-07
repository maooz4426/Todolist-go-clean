package repository

import (
	"github.com/maooz4426/Todolist/domain/entity"
)

// usecase層で使うrepository層のメソッドを定義
// これで依存先を抽象化
// ダックタイピングしてる
type TodoRepositoryer interface {
	Insert(task *entity.Todo) (*entity.Todo, error)
	FindAll() ([]*entity.Todo, error)
	FindById(id string) (*entity.Todo, error)
	Update(task *entity.Todo) (*entity.Todo, error)
	Delete(id string) error
}
