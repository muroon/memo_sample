package repository

import (
	"context"
)

// TransactionRepository Transaction's Repository
type TransactionRepository interface {
	Begin(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
}
