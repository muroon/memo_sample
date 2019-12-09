package memory

import (
	"context"
	"fmt"
	"memo_sample/domain/model"
	"memo_sample/domain/repository"
	"strings"
)

// NewMemoRepository get repository
func NewMemoRepository() repository.MemoRepository {
	return &memoRepository{[]*model.Memo{}}
}

// MemoRepository Memo's Repository Sub
type memoRepository struct {
	memoList []*model.Memo
}

// generateID generate Key
func (m *memoRepository) generateID(ctx context.Context) (int, error) {
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
func (m *memoRepository) Save(ctx context.Context, text string) (*model.Memo, error) {
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
func (m memoRepository) Get(ctx context.Context, id int) (*model.Memo, error) {
	for _, ml := range m.memoList {
		if ml.ID == id {
			return ml, nil
		}
	}
	return nil, fmt.Errorf("Error: %s", "no memo data")
}

// GetAll get all Memo Data
func (m *memoRepository) GetAll(ctx context.Context) ([]*model.Memo, error) {
	return m.memoList, nil
}

// Search search memo by text
func (m *memoRepository) Search(ctx context.Context, text string) ([]*model.Memo, error) {
	list := make([]*model.Memo, 0, len(m.memoList))
	for _, memo := range m.memoList {
		if strings.Contains(memo.Text, text) {
			list = append(list, memo)
		}
	}
	return list, nil
}

// GetAllByIDs get all Memo Data by ID
func (m *memoRepository) GetAllByIDs(ctx context.Context, ids []int) ([]*model.Memo, error) {
	list := make([]*model.Memo, 0, len(m.memoList))
	for _, memo := range m.memoList {
		for _, id := range ids {
			if memo.ID == id {
				list = append(list, memo)
			}
		}
	}
	return list, nil
}
