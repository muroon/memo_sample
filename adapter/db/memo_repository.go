package db

import (
	"context"
	"fmt"
	"memo_sample/domain/model"
	"database/sql"
)

// MemoRepository Memo's Repository Sub
type MemoRepository struct {
	DB *sql.DB
	tx *sql.Tx
}

// ContextKey key for transaction context
type ContextKey string
const (
	txKey = "db.transaction"
)

// Begin begin transaction
func (m *MemoRepository) Begin(ctx context.Context) (context.Context, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}
	m.tx = tx
	ctx = context.WithValue(ctx,ContextKey(txKey),true)
	return ctx, nil
}

// Rollback rollback transaction
func (m *MemoRepository) Rollback(ctx context.Context) (context.Context, error) {
	m.tx.Rollback()
	ctx = context.WithValue(ctx,ContextKey(txKey),false)
	return ctx, nil
}

// Commit commit transaction
func (m *MemoRepository) Commit(ctx context.Context) (context.Context, error) {
	err := m.tx.Commit()
	ctx = context.WithValue(ctx,ContextKey(txKey),false)
	return ctx, err
}

func (m *MemoRepository) isTx(ctx context.Context) bool {
	if txn, ok := ctx.Value(ContextKey(txKey)).(bool); ok {
		return txn
	}
	return false
}

// GenerateID generate Key
func (m *MemoRepository) GenerateID(ctx context.Context) (int, error) {
	lastID := 0

	rows, err := m.DB.Query("select id from memo")
    if err != nil {
        return 0, err
    }
	
	for rows.Next() {
		err := rows.Scan(&lastID)
		if err != nil {
			return 0, err
		}
	}
	return lastID + 1, nil
}

// Save save Memo Data
func (m *MemoRepository) Save(ctx context.Context, memo *model.Memo) error {
	var err error
	if m.isTx(ctx) {
		_, err = m.tx.Exec("insert into memo(text) values(?)", memo.Text)
	} else {
		_, err = m.DB.Exec("insert into memo(text) values(?)", memo.Text)
	}

	return err
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
	rows, err := m.DB.Query("select * from memo")
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