package db

import (
	"context"
	"fmt"
	"memo_sample/domain/model"
	"database/sql"
)

// MemoRepository Memo's Repository Sub
type MemoRepository struct {
	memoList []*model.Memo
	DB *sql.DB
}

// GenerateID generate Key
func (m *MemoRepository) GenerateID(ctx context.Context) (int, error) {
	lastID := 0

	rows, err := m.DB.Query("select id from memo")
    if err != nil {
        panic(err)
    }
	
	for rows.Next() {
		err := rows.Scan(&lastID)
		if err != nil {
			panic(err)
		}
	}
	return lastID + 1, nil
}

// Save save Memo Data
func (m *MemoRepository) Save(ctx context.Context, memo *model.Memo) error {	
	_, err := m.DB.Exec("insert into memo(text) values(?)", memo.Text)
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
        panic(err)
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