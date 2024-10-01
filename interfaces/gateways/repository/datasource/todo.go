package datasource

import (
	"github.com/maooz4426/Todolist/domain/entity"
	"gorm.io/gorm"
	"strconv"
)

type TodoRepository struct {
	db *gorm.DB
}

// インスタンスメソッド
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (m *TodoRepository) Insert(task *entity.Todo) (*entity.Todo, error) {
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

func (m *TodoRepository) FindAll() ([]*entity.Todo, error) {
	var todos []*entity.Todo

	result := m.db.Find(&todos)
	if result.Error != nil {
		return []*entity.Todo{}, result.Error
	}

	return todos, nil
}

func (m *TodoRepository) FindById(id string) (*entity.Todo, error) {
	var todo entity.Todo

	searchId, err := strconv.Atoi(id)
	if err != nil {
		return &entity.Todo{}, err
	}

	result := m.db.First(&todo, "id = ?", searchId)

	if result.Error != nil {
		return &entity.Todo{}, result.Error
	}

	return &todo, nil
}

func (m *TodoRepository) Update(task *entity.Todo) (*entity.Todo, error) {
	m.db.Save(task)

	return task, nil
}
