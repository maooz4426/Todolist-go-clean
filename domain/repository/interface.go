package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain"
)

// usecase層で使うrepository層のメソッドを定義
// これで依存先を抽象化
type TodoRepositoryer interface {
	InsertTodo(ctx echo.Context, task *domain.Todo) (*domain.Todo, error)
}
