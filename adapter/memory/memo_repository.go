package memory

import (
	"context"
	"fmt"
	"memo_sample/domain/model"
)

// MemoRepository Memo's Repository Sub
type MemoRepository struct {
	memoList []*model.Memo
}

// GenerateID generate Key
func (m *MemoRepository) GenerateID(ctx context.Context) (int, error) {
	const initID int = 1

	if len(m.memoList) == 0 {
		return initID, nil
	}

	var lm = m.memoList[len(m.memoList) - 1]
	if lm == nil {
		return initID, nil
	}
	var id = lm.ID + 1
	return id, nil
}

// Save save Memo Data
func (m *MemoRepository) Save(ctx context.Context, memo *model.Memo) error {	
	m.memoList = append(m.memoList, memo)
	return nil
}

// Find get Memo Data by ID
func (m MemoRepository) Find(ctx context.Context, id int) (*model.Memo, error) {
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