package db

import (
	"context"
	"database/sql"
	"fmt"
	"memo_sample/domain/model"
)

// NewDBMemoRepository get repository
func NewDBMemoRepository(db *sql.DB, tx *sql.Tx) *MemoRepository {
	return &MemoRepository{db: db, tx: tx}
}

// MemoRepository Memo's Repository Sub
type MemoRepository struct {
	db *sql.DB
	tx *sql.Tx
}

// ContextKey key for transaction context
type ContextKey string

const (
	txKey = "db.transaction"
)

// Begin begin transaction
func (m *MemoRepository) Begin(ctx context.Context) (context.Context, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	m.tx = tx
	ctx = context.WithValue(ctx, ContextKey(txKey), true)
	return ctx, nil
}

// Rollback rollback transaction
func (m *MemoRepository) Rollback(ctx context.Context) (context.Context, error) {
	m.tx.Rollback()
	ctx = context.WithValue(ctx, ContextKey(txKey), false)
	return ctx, nil
}

// Commit commit transaction
func (m *MemoRepository) Commit(ctx context.Context) (context.Context, error) {
	err := m.tx.Commit()
	ctx = context.WithValue(ctx, ContextKey(txKey), false)
	return ctx, err
}

func (m *MemoRepository) isTx(ctx context.Context) bool {
	if txn, ok := ctx.Value(ContextKey(txKey)).(bool); ok {
		return txn
	}
	return false
}

// Save save Memo Data
func (m *MemoRepository) Save(ctx context.Context, text string) (*model.Memo, error) {
	var err error
	var res sql.Result
	if m.isTx(ctx) {
		res, err = m.tx.Exec("insert into memo(text) values(?)", text)

	} else {
		res, err = m.db.Exec("insert into memo(text) values(?)", text)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return m.Find(ctx, int(id))
}

// Find get Memo Data by ID
func (m MemoRepository) Find(ctx context.Context, id int) (*model.Memo, error) {
	list, err := m.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, ml := range list {
		if ml.ID == id {
			return ml, nil
		}
	}
	return nil, fmt.Errorf("Error: %s", "no memo data")
}

// GetAll get all Memo Data
func (m *MemoRepository) GetAll(ctx context.Context) ([]*model.Memo, error) {
	var rows *sql.Rows
	var err error
	if m.isTx(ctx) {
		rows, err = m.tx.Query("select * from memo")

	} else {
		rows, err = m.db.Query("select * from memo")
	}

	if err != nil {
		return nil, err
	}

	list := []*model.Memo{}
	for rows.Next() {
		mem := &model.Memo{}
		err := rows.Scan(&mem.ID, &mem.Text)
		if err != nil {
			panic(err)
		}
		list = append(list, mem)
	}

	return list, nil
}
