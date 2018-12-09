package usecase

import (
	"context"
	"memo_sample/domain/model"
)

// Presenter presenter interface
type Presenter interface {
	ViewError(ctx context.Context, err error)
	ViewMemo(ctx context.Context, md *model.Memo)
	ViewMemoList(ctx context.Context, list []*model.Memo)
	ViewTag(ctx context.Context, md *model.Tag)
	ViewTagList(ctx context.Context, list []*model.Tag)
	ViewPostMemoAndTagsResult(ctx context.Context, memo *model.Memo, tags []*model.Tag)
	ViewSearchTagsAndMemosResult(ctx context.Context, memos []*model.Memo, tags []*model.Tag)
}
