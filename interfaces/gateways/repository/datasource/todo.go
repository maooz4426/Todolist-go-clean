package datasource

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain/entity"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

// インスタンスメソッド
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (m *TodoRepository) InsertTodo(ctx echo.Context, task *entity.Todo) (*entity.Todo, error) {
	todo := entity.Todo{
		Task:     task.Task,
		Done:     false,
		Deadline: task.Deadline,
	}

	result := m.db.Create(&todo)

	if result.Error != nil {
		return &entity.Todo{}, result.Error
	}

	return &todo, nil
}
