package usecase

import (
	"context"
	"fmt"
	"memo_sample/domain/model"
	"memo_sample/domain/repository"
	"memo_sample/infra/error"
	"memo_sample/usecase/input"
	"net/http"
)

// Memo memo related interface
type Memo interface {
	ValidatePost(ipt input.PostMemo) error
	Post(ctx context.Context, ipt input.PostMemo) (int, error)
	ValidateGet(ipt input.GetMemo) error
	GetMemo(ctx context.Context, ipt input.GetMemo) (*model.Memo, error)
	GetAllMemoList(ctx context.Context) ([]*model.Memo, error)
	ValidatePostMemoAndTags(ipt input.PostMemoAndTags) error
	PostMemoAndTags(ctx context.Context, ipt input.PostMemoAndTags) (*model.Memo, []*model.Tag, error)
	GetTagsByMemo(ctx context.Context, ipt input.GetTagsByMemo) ([]*model.Tag, error)
	SearchTagsAndMemos(ctx context.Context, ipt input.SearchTagsAndMemos) ([]*model.Memo, []*model.Tag, error)
}

// NewMemo generate memo instance
func NewMemo(
	transactionRepository repository.TransactionRepository,
	memoRepository repository.MemoRepository,
	tagRepository repository.TagRepository,
	errm apperror.ErrorManager,
) Memo {
	return memo{
		transactionRepository,
		memoRepository,
		tagRepository,
		errm,
	}
}

type memo struct {
	transactionRepository repository.TransactionRepository
	memoRepository        repository.MemoRepository
	tagRepository         repository.TagRepository
	errm                  apperror.ErrorManager
}

func (m memo) ValidatePost(ipt input.PostMemo) error {
	if ipt.Text == "" {
		err := fmt.Errorf("text parameter is invalid. %s", ipt.Text)
		return m.errm.Wrap(
			err,
			http.StatusBadRequest,
		)
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
		err := fmt.Errorf("ID parameter is invalid. %d", ipt.ID)
		return m.errm.Wrap(
			err,
			http.StatusBadRequest,
		)
	}

	return nil
}

func (m memo) GetMemo(ctx context.Context, ipt input.GetMemo) (*model.Memo, error) {
	me, err := m.memoRepository.Get(ctx, ipt.ID)
	if err != nil {
		return nil, err
	}

	return me, nil
}

func (m memo) GetAllMemoList(ctx context.Context) ([]*model.Memo, error) {
	list, err := m.memoRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (m memo) ValidatePostMemoAndTags(ipt input.PostMemoAndTags) error {
	if ipt.MemoText == "" {
		err := fmt.Errorf("text parameter(MemoText) is invalid. %s", ipt.MemoText)
		return m.errm.Wrap(
			err,
			http.StatusBadRequest,
		)
	}

	for _, title := range ipt.TagTitles {
		if title == "" {
			err := fmt.Errorf("text parameter(TagTitles) is invalid. %s", title)
			return m.errm.Wrap(
				err,
				http.StatusBadRequest,
			)
		}
	}

	return nil
}

func (m memo) PostMemoAndTags(ctx context.Context, ipt input.PostMemoAndTags) (*model.Memo, []*model.Tag, error) {
	tags := []*model.Tag{}

	ctx, err := m.transactionRepository.Begin(ctx)
	if err != nil {
		m.transactionRepository.Rollback(ctx)
		return nil, nil, err
	}

	// Memo
	mo, err := m.memoRepository.Save(ctx, ipt.MemoText)
	if err != nil {
		m.transactionRepository.Rollback(ctx)
		return nil, nil, err
	}

	for _, title := range ipt.TagTitles {
		// Tag
		tg, err := m.tagRepository.Save(ctx, title)
		if err != nil {
			m.transactionRepository.Rollback(ctx)
			return nil, nil, err
		}
		tags = append(tags, tg)

		// MemoTag
		err = m.tagRepository.SaveTagAndMemo(ctx, tg.ID, mo.ID)
		if err != nil {
			m.transactionRepository.Rollback(ctx)
			return nil, nil, err
		}
	}

	m.transactionRepository.Commit(ctx)

	return mo, tags, nil
}

func (m memo) GetTagsByMemo(ctx context.Context, ipt input.GetTagsByMemo) ([]*model.Tag, error) {
	return m.tagRepository.GetAllByMemoID(ctx, ipt.ID)
}

func (m memo) SearchTagsAndMemos(ctx context.Context, ipt input.SearchTagsAndMemos) ([]*model.Memo, []*model.Tag, error) {
	mIDs, err := m.tagRepository.SearchMemoIDsByTitle(ctx, ipt.TagTitle)
	if err != nil {
		return nil, nil, err
	}

	memos, err := m.memoRepository.GetAllByIDs(ctx, mIDs)
	if err != nil {
		return nil, nil, err
	}

	tags := []*model.Tag{}
	for _, mID := range mIDs {
		tgs, err := m.tagRepository.GetAllByMemoID(ctx, mID)
		if err != nil {
			return nil, nil, err
		}
		for _, tg := range tgs {
			tags = append(tgs, tg)
		}
	}

	return memos, tags, nil
}
