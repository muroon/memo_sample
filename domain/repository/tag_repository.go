package repository

import (
	"context"
	"memo_sample/domain/model"
)

// TagRepository Tag's Repository
type TagRepository interface {
	Begin(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
	Save(ctx context.Context, title string) (*model.Tag, context.Context, error)
	Get(ctx context.Context, id int) (*model.Tag, context.Context, error)
	GetAll(ctx context.Context) ([]*model.Tag, context.Context, error)
	Search(ctx context.Context, title string) ([]*model.Tag, context.Context, error)
	SaveTagAndMemo(ctx context.Context, tagID int, memoID int) (context.Context, error)
	GetAllByMemoID(ctx context.Context, id int) ([]*model.Tag, context.Context, error)
	SearchMemoIDsByTitle(ctx context.Context, title string) ([]int, context.Context, error)
}
