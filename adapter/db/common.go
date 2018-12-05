package db

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

var db *sql.DB
var txMap map[string]*sql.Tx
var stmt *sql.Stmt

// setDB db setting
func setDB(d *sql.DB) {
	if db == nil {
		db = d
	}
}

// ContextKey key for transaction context
type ContextKey string

const (
	txIsKey = "db.transaction"
	txKey   = "db.transaction.key"
)

func init() {

	if txMap == nil {
		txMap = map[string]*sql.Tx{}
	}
}

// getDB
func getDB(ctx context.Context) *sql.DB {
	return db
}

func getTx(ctx context.Context) *sql.Tx {
	key := getTxKey(ctx)
	fmt.Println("getTx key:" + key) // TODO:

	return txMap[key]
}

func setTx(ctx context.Context, t *sql.Tx) context.Context {
	key := generateNewKey()
	txMap[key] = t
	ctx = setTxKey(ctx, key)

	// Transaction開始フラグ
	ctx = context.WithValue(ctx, ContextKey(txIsKey), true)

	return ctx
}

func deleteTx(ctx context.Context) context.Context {
	txKey := getTxKey(ctx)
	if _, ok := txMap[txKey]; ok {
		delete(txMap, txKey)
	}

	// Transaction開始フラグ
	return context.WithValue(ctx, ContextKey(txIsKey), false)
}

// begin begin transaction
func begin(ctx context.Context) (context.Context, error) {
	t, err := getDB(ctx).BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Transaction関連の設定
	ctx = setTx(ctx, t)

	return ctx, nil
}

// rollback rollback transaction
func rollback(ctx context.Context) (context.Context, error) {
	getTx(ctx).Rollback()

	// Transaction関連削除
	ctx = deleteTx(ctx)

	return ctx, nil
}

// commit commit transaction
func commit(ctx context.Context) (context.Context, error) {
	err := getTx(ctx).Commit()

	// Transaction関連削除
	ctx = deleteTx(ctx)

	return ctx, err
}

// prepare
func prepare(ctx context.Context, query string) (*sql.Stmt, context.Context, error) {
	//var st *sql.Stmt
	var err error
	if isTx(ctx) {
		stmt, err = getTx(ctx).PrepareContext(ctx, query)
	} else {
		stmt, err = getDB(ctx).PrepareContext(ctx, query)
	}
	if err != nil {
		return stmt, ctx, err
	}

	// TODO:
	fmt.Println("prepare query:" + query)
	fmt.Printf("prepare isTx:%t \n", isTx(ctx))

	return stmt, ctx, nil
}

// isTx is in transaction or not
func isTx(ctx context.Context) bool {
	if txn, ok := ctx.Value(ContextKey(txIsKey)).(bool); ok {
		return txn
	}
	return false
}

// setTxKey
func setTxKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ContextKey(txKey), value)
}

// getTxKey
func getTxKey(ctx context.Context) string {
	return getKey(ctx, ContextKey(txKey))
}

// getKey get key
func getKey(ctx context.Context, ctxKey ContextKey) string {
	key, _ := ctx.Value(ctxKey).(string)
	return key
}

// generateNewKey generate key
func generateNewKey() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Int())
}
