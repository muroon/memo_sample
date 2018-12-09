package usecase

import (
	"context"
	"memo_sample/usecase/input"
)

// NewInteractor new Interactor
func NewInteractor(
	pre Presenter,
	me Memo,
) Interactor {
	return Interactor{pre, me}
}

// Interactor usecase interactor
type Interactor struct {
	pre  Presenter
	memo Memo
}

// PostMemo post memo
func (i Interactor) PostMemo(ctx context.Context, ipt input.PostMemo) {

	id, err := i.memo.Post(ctx, ipt)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	iptf := &input.GetMemo{ID: id}
	memo, err := i.memo.GetMemo(ctx, *iptf)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	i.pre.ViewMemo(ctx, memo)
}

// GetMemos get all memos
func (i Interactor) GetMemos(ctx context.Context) {

	memos, err := i.memo.GetAllMemoList(ctx)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	i.pre.ViewMemoList(ctx, memos)
}

// PostMemoAndTags save memo and tags
func (i Interactor) PostMemoAndTags(ctx context.Context, ipt input.PostMemoAndTags) {

	memo, tags, err := i.memo.PostMemoAndTags(ctx, ipt)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	i.pre.ViewPostMemoAndTagsResult(ctx, memo, tags)
}

// SearchTagsAndMemos save memo and tags
func (i Interactor) SearchTagsAndMemos(ctx context.Context, ipt input.SearchTagsAndMemos) {

	memos, tags, err := i.memo.SearchTagsAndMemos(ctx, ipt)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	i.pre.ViewSearchTagsAndMemosResult(ctx, memos, tags)
}
