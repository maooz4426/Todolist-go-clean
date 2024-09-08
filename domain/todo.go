package domain

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Task     string
	Deadline time.Time
}
