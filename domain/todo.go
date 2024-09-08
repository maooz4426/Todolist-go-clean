package domain

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Task     string
	Done     bool
	Deadline time.Time
}
