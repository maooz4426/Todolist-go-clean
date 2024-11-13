package persistence

import (
	"context"
	"gorm.io/gorm"
)

type TransactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{db}
}

func (t *TransactionManager) RunInTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx := t.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := f(ctx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
