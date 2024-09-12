package usecases

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain/entity"
)

type TodoUseCaser interface {
	Create(ctx echo.Context, task *entity.Todo) error
}
