package memory

import (
	"context"
	"memo_sample/domain/repository"
)

// NewTransactionRepository get repository
func NewTransactionRepository() repository.TransactionRepository {
	return transactionRepository{}
}

// transactionRepository transaction Repository Sub
type transactionRepository struct{}

func (t transactionRepository) Begin(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (t transactionRepository) Rollback(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (t transactionRepository) Commit(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
