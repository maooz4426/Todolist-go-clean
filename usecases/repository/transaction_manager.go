package repository

import "context"

type ITransactionManager interface {
	RunInTx(ctx context.Context, f func(ctx context.Context) error) error
}
