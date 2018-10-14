package repository

import (
	"context"
	"memo_sample/domain/model"
)

// MemoRepository Memo's Repository
type MemoRepository interface {
	GenerateID(ctx context.Context) (int, error)
    Save(ctx context.Context, memo *model.Memo) error
	Find(ctx context.Context, id int) (*model.Memo, error)
	GetAll(ctx context.Context) ([]*model.Memo, error)
}