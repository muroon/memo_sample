package repository

import (
	"context"
	"memo_sample/domain/model"
)

// MemoRepository Memo's Repository
type MemoRepository interface {
	Begin(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
	Save(ctx context.Context, text string) (*model.Memo, context.Context, error)
	Get(ctx context.Context, id int) (*model.Memo, context.Context, error)
	GetAll(ctx context.Context) ([]*model.Memo, context.Context, error)
	Search(ctx context.Context, text string) ([]*model.Memo, context.Context, error)
	GetAllByIDs(ctx context.Context, ids []int) ([]*model.Memo, context.Context, error)
}
