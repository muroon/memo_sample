package usecase

import (
	"context"
	"fmt"
	"memo_sample/domain/model"
	"memo_sample/domain/repository"
	"memo_sample/usecase/input"
	"memo_sample/usecase/output/json"
)

// Memo memo related interface
type Memo interface {
	ValidatePost(ipt input.PostMemo) error
	Post(ctx context.Context, ipt input.PostMemo) (int, error)
	ValidateGet(ipt input.GetMemo) error
	Get(ctx context.Context, ipt input.GetMemo) (*model.Memo, error)
	GetAll(ctx context.Context) ([]*model.Memo, error)
	GetJSON(ctx context.Context, ipt input.GetMemo) (*json.Memo, error)
	GetAllJSON(ctx context.Context) ([]*json.Memo, error)
}

// NewMemo generate memo instance
func NewMemo(
	memoRepository repository.MemoRepository,
	tagRepository repository.TagRepository,
) Memo {
	return memo{
		memoRepository,
		tagRepository,
	}
}

type memo struct {
	memoRepository repository.MemoRepository
	tagRepository  repository.TagRepository
}

func (m memo) ValidatePost(ipt input.PostMemo) error {
	if ipt.Text == "" {
		return fmt.Errorf("text parameter is invalid. %s", ipt.Text)
	}

	return nil
}

func (m memo) Post(ctx context.Context, ipt input.PostMemo) (int, error) {
	mo, err := m.memoRepository.Save(ctx, ipt.Text)
	if err != nil {
		return 0, err
	}
	return mo.ID, err
}

func (m memo) ValidateGet(ipt input.GetMemo) error {
	if ipt.ID <= 0 {
		return fmt.Errorf("ID parameter is invalid. %d", ipt.ID)
	}

	return nil
}

func (m memo) Get(ctx context.Context, ipt input.GetMemo) (*model.Memo, error) {
	return m.memoRepository.Get(ctx, ipt.ID)
}

func (m memo) GetAll(ctx context.Context) ([]*model.Memo, error) {
	return m.memoRepository.GetAll(ctx)
}

func (m memo) changeJSON(md *model.Memo) *json.Memo {
	mj := &json.Memo{
		ID:   md.ID,
		Text: md.Text,
	}
	return mj
}

func (m memo) GetJSON(ctx context.Context, ipt input.GetMemo) (*json.Memo, error) {
	md, err := m.Get(ctx, ipt)
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
