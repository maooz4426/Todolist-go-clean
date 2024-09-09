package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

// インスタンスメソッド
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (m *TodoRepository) InsertTodo(ctx echo.Context, task *domain.Todo) (*domain.Todo, error) {
	todo := domain.Todo{
		Task:     task.Task,
		Done:     false,
		Deadline: task.Deadline,
	}

	result := m.db.Create(&todo)

	if result.Error != nil {
		return &domain.Todo{}, result.Error
	}

	return &todo, nil
}
