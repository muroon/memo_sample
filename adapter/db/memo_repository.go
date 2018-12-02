package db

import (
	"context"
	"database/sql"
	"fmt"
	"memo_sample/domain/model"
)

// NewMemoRepository get repository
func NewMemoRepository(db *sql.DB) *MemoRepository {
	setDB(db)
	return &MemoRepository{}
}

// MemoRepository Memo's Repository Sub
type MemoRepository struct{}

// Begin begin transaction
func (m *MemoRepository) Begin(ctx context.Context) (context.Context, error) {
	return begin(ctx)
}

// Rollback rollback transaction
func (m *MemoRepository) Rollback(ctx context.Context) (context.Context, error) {
	return rollback(ctx)
}

// Commit commit transaction
func (m *MemoRepository) Commit(ctx context.Context) (context.Context, error) {
	return commit(ctx)
}

// Save save Memo Data
func (m *MemoRepository) Save(ctx context.Context, text string) (*model.Memo, error) {
	var err error
	var res sql.Result
	if isTx(ctx) {
		res, err = tx.Exec("insert into memo(text) values(?)", text)

	} else {
		res, err = db.Exec("insert into memo(text) values(?)", text)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return m.Get(ctx, int(id))
}

// Get get Memo Data by ID
func (m MemoRepository) Get(ctx context.Context, id int) (*model.Memo, error) {
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
	if isTx(ctx) {
		rows, err = tx.Query("select * from memo")

	} else {
		rows, err = db.Query("select * from memo")
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
