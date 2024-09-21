package port

import (
	"github.com/maooz4426/Todolist/domain/entity"
)

type TodoUseCaser interface {
	Create(task *entity.Todo) (*entity.Todo, error)
	FindAll() ([]*entity.Todo, error)
	FindById(id string) (*entity.Todo, error)
}
