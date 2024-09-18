package port

import (
	"github.com/maooz4426/Todolist/domain/dto"
	"github.com/maooz4426/Todolist/domain/entity"
)

type TodoUseCaser interface {
	Create(task *entity.Todo) (*dto.CreateResponse, error)
}
