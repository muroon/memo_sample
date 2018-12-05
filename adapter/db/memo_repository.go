package db

import (
	"context"
	"database/sql"
	"memo_sample/domain/model"
	"strconv"
	"strings"
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
func (m *MemoRepository) Save(ctx context.Context, text string) (*model.Memo, context.Context, error) {
	var err error
	var res sql.Result
	query := "insert into memo(text) values(?)"
	stmt, ctx, err := prepare(ctx, query)
	if err != nil {
		return nil, ctx, err
	}

	res, err = stmt.ExecContext(ctx, text)
	if err != nil {
		return nil, ctx, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, ctx, err
	}

	return m.Get(ctx, int(id))
}

// Get get Memo Data by ID
func (m MemoRepository) Get(ctx context.Context, id int) (*model.Memo, context.Context, error) {
	mem := &model.Memo{}
	var err error
	query := "select * from memo where id = ?"
	stmt, ctx, err := prepare(ctx, query)
	if err != nil {
		return nil, ctx, err
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&mem.ID, &mem.Text)
	if err != nil {
		return nil, ctx, err
	}

	return mem, ctx, err
}

// GetAll get all Memo Data
func (m *MemoRepository) GetAll(ctx context.Context) ([]*model.Memo, context.Context, error) {
	var rows *sql.Rows
	var err error
	query := "select * from memo"
	stmt, ctx, err := prepare(ctx, query)
	if err != nil {
		return nil, ctx, err
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, ctx, err
	}

	return m.getModelList(ctx, rows)
}

// Search search memo by text
func (m *MemoRepository) Search(ctx context.Context, text string) ([]*model.Memo, context.Context, error) {
	var rows *sql.Rows
	var err error
	query := "select * from memo where text like '%" + text + "%'"
	stmt, ctx, err := prepare(ctx, query)
	if err != nil {
		return nil, ctx, err
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, ctx, err
	}

	return m.getModelList(ctx, rows)
}

// GetAllByIDs get all Memo Data by ID
func (m *MemoRepository) GetAllByIDs(ctx context.Context, ids []int) ([]*model.Memo, context.Context, error) {
	idvs := []string{}
	for _, id := range ids {
		idvs = append(idvs, strconv.Itoa(id))
	}

	query := "select * from memo where id in (" + strings.Join(idvs, ",") + ")"

	var rows *sql.Rows
	var err error
	stmt, ctx, err := prepare(ctx, query)
	if err != nil {
		return nil, ctx, err
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, ctx, err
	}

	return m.getModelList(ctx, rows)
}

// getModelList get model list
func (m *MemoRepository) getModelList(ctx context.Context, rows *sql.Rows) ([]*model.Memo, context.Context, error) {
	list := []*model.Memo{}
	for rows.Next() {
		mem := &model.Memo{}
		err := rows.Scan(&mem.ID, &mem.Text)
		if err != nil {
			return list, ctx, err
		}
		list = append(list, mem)
	}

	return list, ctx, nil
}
