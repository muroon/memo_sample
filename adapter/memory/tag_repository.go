package memory

import (
	"context"
	"fmt"
	"memo_sample/domain/model"
	"memo_sample/domain/repository"
	"strings"
)

// NewTagRepository get repository
func NewTagRepository() repository.TagRepository {
	return &tagRepository{tagList: []*model.Tag{}, tagMemoMap: map[int]int{}}
}

// tagRepository Tag's Repository Sub
type tagRepository struct {
	tagList    []*model.Tag
	tagMemoMap map[int]int
}

// generateID generate Key
func (m *tagRepository) generateID(ctx context.Context) (int, error) {
	const initID int = 1

	if len(m.tagList) == 0 {
		return initID, nil
	}

	var lm = m.tagList[len(m.tagList)-1]
	if lm == nil {
		return initID, nil
	}
	var id = lm.ID + 1
	return id, nil
}

// Save save Tag Data
func (m *tagRepository) Save(ctx context.Context, title string) (*model.Tag, error) {
	id, err := m.generateID(ctx)
	if err != nil {
		return nil, err
	}

	tag := &model.Tag{
		ID:    id,
		Title: title,
	}

	m.tagList = append(m.tagList, tag)
	return tag, nil
}

// Get get Tag Data by ID
func (m tagRepository) Get(ctx context.Context, id int) (*model.Tag, error) {
	for _, ml := range m.tagList {
		if ml.ID == id {
			return ml, nil
		}
	}
	return nil, fmt.Errorf("Error: %s", "no tag data")
}

// GetAll get all Tag Data
func (m *tagRepository) GetAll(ctx context.Context) ([]*model.Tag, error) {
	return m.tagList, nil
}

// Search search tag by text
func (m *tagRepository) Search(ctx context.Context, title string) ([]*model.Tag, error) {
	list := []*model.Tag{}
	for _, tag := range m.tagList {
		if strings.Index(tag.Title, title) != -1 {
			list = append(list, tag)
		}
	}
	return list, nil
}

// SaveTagAndMemo save tag and memo link
func (m *tagRepository) SaveTagAndMemo(ctx context.Context, tagID int, memoID int) error {
	m.tagMemoMap[tagID] = memoID

	return nil
}

// GetAllByMemoID get all Tag Data By MemoID
func (m *tagRepository) GetAllByMemoID(ctx context.Context, id int) ([]*model.Tag, error) {
	list := []*model.Tag{}

	for i, v := range m.tagMemoMap {
		if v != id {
			continue
		}

		for _, tag := range m.tagList {
			if i != tag.ID {
				continue
			}
			list = append(list, tag)
		}
	}

	return list, nil
}

// SearchMemoIDsByTitle search memo ids by tag's title
func (m *tagRepository) SearchMemoIDsByTitle(ctx context.Context, title string) ([]int, error) {
	memoIDs := []int{}

	list, err := m.Search(ctx, title)
	if err != nil {
		return memoIDs, err
	}

	for _, tag := range list {
		for tagID, memoID := range m.tagMemoMap {
			if tag.ID != tagID {
				continue
			}
			memoIDs = append(memoIDs, memoID)
		}
	}

	return memoIDs, nil
}
