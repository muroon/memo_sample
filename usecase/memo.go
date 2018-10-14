package usecase

import (
	"context"
	"memo_sample/domain/repository"
	"memo_sample/domain/model"
	"memo_sample/output/json"
)

// Memo memo related interface
type Memo interface {
	Post(ctx context.Context, text string) (int, error)
	Find(ctx context.Context, id int) (*model.Memo, error)
	GetAll(ctx context.Context) ([]*model.Memo, error)
	FindJSON(ctx context.Context, id int) (*json.Memo, error)
	GetAllJSON(ctx context.Context) ([]*json.Memo, error)
}

// NewMemo generate memo instance
func NewMemo(
	memoRepository repository.MemoRepository,
) Memo {
	return memo{
		memoRepository,
	}
}

type memo struct {
	memoRepository repository.MemoRepository
}

func (m memo) Post(ctx context.Context, text string) (int, error) {
	id, err := m.memoRepository.GenerateID(ctx)
	if err != nil {
		return 0, err
	}

	mo := &model.Memo{
		ID: id,
		Text: text,
	}

	err = m.memoRepository.Save(ctx, mo)
	return id, err
}

func (m memo) Find(ctx context.Context, id int) (*model.Memo, error) {
	return m.memoRepository.Find(ctx, id)
}

func (m memo) GetAll(ctx context.Context) ([]*model.Memo, error) {
	return m.memoRepository.GetAll(ctx)
}

func (m memo) changeJSON(md *model.Memo) *json.Memo {
	mj := &json.Memo {
		ID: md.ID,
		Text: md.Text,
	}
	return mj
}

func (m memo) FindJSON(ctx context.Context, id int) (*json.Memo, error) {
	md, err := m.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return m.changeJSON(md), nil
}

func (m memo) GetAllJSON(ctx context.Context) ([]*json.Memo, error) {
	list, err := m.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	listJSON := []*json.Memo{}
	for _, v := range list {
		listJSON = append(listJSON, m.changeJSON(v))
	}
	return listJSON, nil
}