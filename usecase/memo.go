package usecase

import (
	"context"
	"fmt"
	"memo_sample/domain/repository"
	"memo_sample/domain/model"
	"memo_sample/usecase/output/json"
	"memo_sample/usecase/input"
)

// Memo memo related interface
type Memo interface {
	ValidatePost(ipt input.PostMemo) error
	Post(ctx context.Context, ipt input.PostMemo) (int, error)
	ValidateFind(ipt input.FindMemo) error
	Find(ctx context.Context, ipt input.FindMemo) (*model.Memo, error)
	GetAll(ctx context.Context) ([]*model.Memo, error)
	FindJSON(ctx context.Context, ipt input.FindMemo) (*json.Memo, error)
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

func (m memo) ValidatePost(ipt input.PostMemo) error {
	if ipt.Text == "" {
		return fmt.Errorf("text parameter is invalid. %s", ipt.Text)
	}

	return nil
}

func (m memo) Post(ctx context.Context, ipt input.PostMemo) (int, error) {
	id, err := m.memoRepository.GenerateID(ctx)
	if err != nil {
		return 0, err
	}

	mo := &model.Memo{
		ID: id,
		Text: ipt.Text,
	}

	err = m.memoRepository.Save(ctx, mo)
	return id, err
}

func (m memo) ValidateFind(ipt input.FindMemo) error {
	if ipt.ID <= 0 {
		return fmt.Errorf("ID parameter is invalid. %d", ipt.ID)
	}

	return nil
}

func (m memo) Find(ctx context.Context, ipt input.FindMemo) (*model.Memo, error) {
	return m.memoRepository.Find(ctx, ipt.ID)
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

func (m memo) FindJSON(ctx context.Context, ipt input.FindMemo) (*json.Memo, error) {
	md, err := m.Find(ctx, ipt)
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