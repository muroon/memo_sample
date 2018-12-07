package db

import (
	"context"
	"memo_sample/domain/repository"
	"memo_sample/infra"
)

// NewTransactionRepository get repository
func NewTransactionRepository(d *infra.DBM) repository.TransactionRepository {
	setDBM(d)
	return transactionRepository{}
}

// transactionRepository transaction Repository Sub
type transactionRepository struct{}

// Begin begin transaction
func (t transactionRepository) Begin(ctx context.Context) (context.Context, error) {
	return begin(ctx)
}

// Rollback rollback transaction
func (t transactionRepository) Rollback(ctx context.Context) (context.Context, error) {
	return rollback(ctx)
}

// Commit commit transaction
func (t transactionRepository) Commit(ctx context.Context) (context.Context, error) {
	return commit(ctx)
}
