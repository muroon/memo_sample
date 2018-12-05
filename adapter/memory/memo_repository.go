package memory

import (
	"context"
	"fmt"
	"memo_sample/domain/model"
	"strings"
)

// NewMemoRepository get repository
func NewMemoRepository() *MemoRepository {
	return &MemoRepository{}
}

// MemoRepository Memo's Repository Sub
type MemoRepository struct {
	memoList []*model.Memo
}

// Begin begin transaction
func (m *MemoRepository) Begin(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// Rollback rollback transaction
func (m *MemoRepository) Rollback(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// Commit commit transaction
func (m *MemoRepository) Commit(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// generateID generate Key
func (m *MemoRepository) generateID(ctx context.Context) (int, error) {
	const initID int = 1

	if len(m.memoList) == 0 {
		return initID, nil
	}

	var lm = m.memoList[len(m.memoList)-1]
	if lm == nil {
		return initID, nil
	}
	var id = lm.ID + 1
	return id, nil
}

// Save save Memo Data
func (m *MemoRepository) Save(ctx context.Context, text string) (*model.Memo, context.Context, error) {
	id, err := m.generateID(ctx)
	if err != nil {
		return nil, ctx, err
	}

	memo := &model.Memo{
		ID:   id,
		Text: text,
	}

	m.memoList = append(m.memoList, memo)
	return memo, ctx, nil
}

// Get get Memo Data by ID
func (m MemoRepository) Get(ctx context.Context, id int) (*model.Memo, context.Context, error) {
	for _, ml := range m.memoList {
		if ml.ID == id {
			return ml, ctx, nil
		}
	}
	return nil, ctx, fmt.Errorf("Error: %s", "no memo data")
}

// GetAll get all Memo Data
func (m *MemoRepository) GetAll(ctx context.Context) ([]*model.Memo, context.Context, error) {
	return m.memoList, ctx, nil
}

// Search search memo by text
func (m *MemoRepository) Search(ctx context.Context, text string) ([]*model.Memo, context.Context, error) {
	list := []*model.Memo{}
	for _, memo := range m.memoList {
		if strings.Index(memo.Text, text) != -1 {
			list = append(list, memo)
		}
	}
	return list, ctx, nil
}

// GetAllByIDs get all Memo Data by ID
func (m *MemoRepository) GetAllByIDs(ctx context.Context, ids []int) ([]*model.Memo, context.Context, error) {
	list := []*model.Memo{}
	for _, memo := range m.memoList {
		for _, id := range ids {
			if memo.ID == id {
				list = append(list, memo)
			}
		}
	}
	return list, ctx, nil
}
