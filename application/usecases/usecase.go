package usecases

import (
	"gorm.io/gorm"
)

type UseCase struct {
	db *gorm.DB
}

func NewUseCase(db *gorm.DB) *UseCase {
	return &UseCase{db: db}
}
