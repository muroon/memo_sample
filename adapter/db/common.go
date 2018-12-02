package db

import (
	"context"
	"database/sql"
)

var db *sql.DB
var tx *sql.Tx

// setDB db setting
func setDB(d *sql.DB) {
	if db == nil {
		db = d
	}
}

// ContextKey key for transaction context
type ContextKey string

const (
	txKey = "db.transaction"
)

// begin begin transaction
func begin(ctx context.Context) (context.Context, error) {
	t, err := db.Begin()
	if err != nil {
		return nil, err
	}
	tx = t
	ctx = context.WithValue(ctx, ContextKey(txKey), true)
	return ctx, nil
}

// rollback rollback transaction
func rollback(ctx context.Context) (context.Context, error) {
	tx.Rollback()
	ctx = context.WithValue(ctx, ContextKey(txKey), false)
	return ctx, nil
}

// commit commit transaction
func commit(ctx context.Context) (context.Context, error) {
	err := tx.Commit()
	ctx = context.WithValue(ctx, ContextKey(txKey), false)
	return ctx, err
}

// isTx is in transaction or not
func isTx(ctx context.Context) bool {
	if txn, ok := ctx.Value(ContextKey(txKey)).(bool); ok {
		return txn
	}
	return false
}
