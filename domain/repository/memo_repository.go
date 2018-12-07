package repository

import (
	"context"
	"memo_sample/domain/model"
)

// MemoRepository Memo's Repository
type MemoRepository interface {
	Save(ctx context.Context, text string) (*model.Memo, error)
	Get(ctx context.Context, id int) (*model.Memo, error)
	GetAll(ctx context.Context) ([]*model.Memo, error)
	Search(ctx context.Context, text string) ([]*model.Memo, error)
	GetAllByIDs(ctx context.Context, ids []int) ([]*model.Memo, error)
}
