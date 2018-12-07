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
func (m *MemoRepository) Save(ctx context.Context, text string) (*model.Memo, error) {
	id, err := m.generateID(ctx)
	if err != nil {
		return nil, err
	}

	memo := &model.Memo{
		ID:   id,
		Text: text,
	}

	m.memoList = append(m.memoList, memo)
	return memo, nil
}

// Get get Memo Data by ID
func (m MemoRepository) Get(ctx context.Context, id int) (*model.Memo, error) {
	for _, ml := range m.memoList {
		if ml.ID == id {
			return ml, nil
		}
	}
	return nil, fmt.Errorf("Error: %s", "no memo data")
}

// GetAll get all Memo Data
func (m *MemoRepository) GetAll(ctx context.Context) ([]*model.Memo, error) {
	return m.memoList, nil
}

// Search search memo by text
func (m *MemoRepository) Search(ctx context.Context, text string) ([]*model.Memo, error) {
	list := []*model.Memo{}
	for _, memo := range m.memoList {
		if strings.Index(memo.Text, text) != -1 {
			list = append(list, memo)
		}
	}
	return list, nil
}

// GetAllByIDs get all Memo Data by ID
func (m *MemoRepository) GetAllByIDs(ctx context.Context, ids []int) ([]*model.Memo, error) {
	list := []*model.Memo{}
	for _, memo := range m.memoList {
		for _, id := range ids {
			if memo.ID == id {
				list = append(list, memo)
			}
		}
	}
	return list, nil
}
