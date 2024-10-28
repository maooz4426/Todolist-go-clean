package entity

import (
	"github.com/maooz4426/Todolist/domain/dto"
	"gorm.io/gorm"
	"time"
)

//type CustomModel struct {
//	ID        string `gorm:"primarykey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}

type Todo struct {
	gorm.Model
	Task     string
	Done     bool
	Deadline time.Time
}

func (task *Todo) ConvertDTO() (*dto.TodoJson, error) {
	format := "2006-01-02"

	timeF := task.Deadline.Format(format)

	res := &dto.TodoJson{
		ID:       task.ID,
		Task:     task.Task,
		Deadline: timeF,
		Done:     task.Done,
	}

	return res, nil
}
