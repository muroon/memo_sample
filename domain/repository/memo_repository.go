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
	Save(ctx context.Context, text string) (*model.Memo, error)
	Get(ctx context.Context, id int) (*model.Memo, error)
	GetAll(ctx context.Context) ([]*model.Memo, error)
}
