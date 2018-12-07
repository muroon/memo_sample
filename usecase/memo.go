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
	GetMemo(ctx context.Context, ipt input.GetMemo) (*json.Memo, error)
	GetAllMemoList(ctx context.Context) ([]*json.Memo, error)
	ValidatePostMemoAndTags(ipt input.PostMemoAndTags) error
	PostMemoAndTags(ctx context.Context, ipt input.PostMemoAndTags) (*json.PostMemoAndTagsResult, error)
	GetTagsByMemo(ctx context.Context, ipt input.GetTagsByMemo) ([]*json.Tag, error)
	SearchTagsAndMemos(ctx context.Context, ipt input.SearchTagsAndMemos) (*json.SearchTagsAndMemosResult, error)
}

// NewMemo generate memo instance
func NewMemo(
	transactionRepository repository.TransactionRepository,
	memoRepository repository.MemoRepository,
	tagRepository repository.TagRepository,
) Memo {
	return memo{
		transactionRepository,
		memoRepository,
		tagRepository,
	}
}

type memo struct {
	transactionRepository repository.TransactionRepository
	memoRepository        repository.MemoRepository
	tagRepository         repository.TagRepository
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

func (m memo) GetMemo(ctx context.Context, ipt input.GetMemo) (*json.Memo, error) {
	me, err := m.memoRepository.Get(ctx, ipt.ID)
	if err != nil {
		return nil, err
	}

	return m.convertMemoJSON(me), nil
}

func (m memo) GetAllMemoList(ctx context.Context) ([]*json.Memo, error) {
	list, err := m.memoRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return m.convertMemoJSONList(list), nil
}

func (m memo) ValidatePostMemoAndTags(ipt input.PostMemoAndTags) error {
	if ipt.MemoText == "" {
		return fmt.Errorf("text parameter(MemoText) is invalid. %s", ipt.MemoText)
	}

	for _, title := range ipt.TagTitles {
		if title == "" {
			return fmt.Errorf("text parameter(TagTitles) is invalid. %s", title)
		}
	}

	return nil
}

func (m memo) PostMemoAndTags(ctx context.Context, ipt input.PostMemoAndTags) (*json.PostMemoAndTagsResult, error) {
	tags := []*model.Tag{}

	ctx, err := m.transactionRepository.Begin(ctx)
	if err != nil {
		m.transactionRepository.Rollback(ctx)
		return nil, err
	}

	// Memo
	mo, err := m.memoRepository.Save(ctx, ipt.MemoText)
	if err != nil {
		m.transactionRepository.Rollback(ctx)
		return nil, err
	}

	for _, title := range ipt.TagTitles {
		// Tag
		tg, err := m.tagRepository.Save(ctx, title)
		if err != nil {
			m.transactionRepository.Rollback(ctx)
			return nil, err
		}
		tags = append(tags, tg)

		// MemoTag
		err = m.tagRepository.SaveTagAndMemo(ctx, tg.ID, mo.ID)
		if err != nil {
			m.transactionRepository.Rollback(ctx)
			return nil, err
		}
	}

	m.transactionRepository.Commit(ctx)

	return &json.PostMemoAndTagsResult{
		Memo: m.convertMemoJSON(mo),
		Tags: m.convertTagJSONList(tags),
	}, nil
}

func (m memo) GetTagsByMemo(ctx context.Context, ipt input.GetTagsByMemo) ([]*json.Tag, error) {
	tags, err := m.tagRepository.GetAllByMemoID(ctx, ipt.ID)
	if err != nil {
		return nil, err
	}
	return m.convertTagJSONList(tags), nil
}

func (m memo) SearchTagsAndMemos(ctx context.Context, ipt input.SearchTagsAndMemos) (*json.SearchTagsAndMemosResult, error) {
	mIDs, err := m.tagRepository.SearchMemoIDsByTitle(ctx, ipt.TagTitle)
	if err != nil {
		return nil, err
	}

	memos, err := m.memoRepository.GetAllByIDs(ctx, mIDs)
	if err != nil {
		return nil, err
	}

	tags := []*model.Tag{}
	for _, mID := range mIDs {
		tgs, err := m.tagRepository.GetAllByMemoID(ctx, mID)
		if err != nil {
			return nil, err
		}
		for _, tg := range tgs {
			tags = append(tgs, tg)
		}
	}

	return &json.SearchTagsAndMemosResult{
		Tags:  m.convertTagJSONList(tags),
		Memos: m.convertMemoJSONList(memos),
	}, nil
}

func (m memo) convertMemoJSON(md *model.Memo) *json.Memo {
	mj := &json.Memo{
		ID:   md.ID,
		Text: md.Text,
	}
	return mj
}

func (m memo) convertMemoJSONList(list []*model.Memo) []*json.Memo {
	listJSON := []*json.Memo{}
	for _, v := range list {
		listJSON = append(listJSON, m.convertMemoJSON(v))
	}
	return listJSON
}

func (m memo) convertTagJSON(md *model.Tag) *json.Tag {
	mj := &json.Tag{
		ID:    md.ID,
		Title: md.Title,
	}
	return mj
}

func (m memo) convertTagJSONList(list []*model.Tag) []*json.Tag {
	listJSON := []*json.Tag{}
	for _, v := range list {
		listJSON = append(listJSON, m.convertTagJSON(v))
	}
	return listJSON
}
