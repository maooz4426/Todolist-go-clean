package repository

import (
	"context"
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

func (m *TodoRepository) Insert(ctx context.Context, task *entity.Todo) (*entity.Todo, error) {
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

func (m *TodoRepository) FindAll(ctx context.Context) ([]*entity.Todo, error) {
	var todos []*entity.Todo

	result := m.db.Find(&todos)
	if result.Error != nil {
		return []*entity.Todo{}, result.Error
	}

	return todos, nil
}

func (m *TodoRepository) FindById(ctx context.Context, id string) (*entity.Todo, error) {
	var todo entity.Todo

	//searchId, err := strconv.Atoi(id)
	//if err != nil {
	//	return &entity.Todo{}, err
	//}

	result := m.db.First(&todo, "id = ?", id)

	if result.Error != nil {
		return &entity.Todo{}, result.Error
	}

	return &todo, nil
}

func (m *TodoRepository) Update(ctx context.Context, task *entity.Todo) (*entity.Todo, error) {
	m.db.Save(task)

	return task, nil
}

func (m *TodoRepository) Delete(ctx context.Context, id string) error {
	m.db.Delete(&entity.Todo{}, id)

	return nil
}
