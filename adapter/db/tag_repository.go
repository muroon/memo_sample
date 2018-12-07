package db

import (
	"context"
	"database/sql"
	"memo_sample/domain/model"
	"memo_sample/domain/repository"
)

// NewTagRepository get repository
func NewTagRepository() repository.TagRepository {
	return tagRepository{}
}

// tagRepository Tag's Repository Sub
type tagRepository struct{}

// Save save Tag Data
func (m tagRepository) Save(ctx context.Context, title string) (*model.Tag, error) {
	var err error
	var res sql.Result
	query := "insert into tag(title) values(?)"
	stmt, err := prepare(ctx, query)
	if err != nil {
		return nil, err
	}

	res, err = stmt.ExecContext(ctx, title)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return m.Get(ctx, int(id))
}

// Get get Tag Data by ID
func (m tagRepository) Get(ctx context.Context, id int) (*model.Tag, error) {
	tag := &model.Tag{}
	var err error
	query := "select * from tag where id = ?"
	stmt, err := prepare(ctx, query)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&tag.ID, &tag.Title)
	if err != nil {
		return nil, err
	}

	return tag, err
}

// GetAll get all Tag Data
func (m tagRepository) GetAll(ctx context.Context) ([]*model.Tag, error) {
	var rows *sql.Rows
	var err error
	query := "select * from tag"
	stmt, err := prepare(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	return m.getModelList(rows)
}

// Search search list by title
func (m tagRepository) Search(ctx context.Context, title string) ([]*model.Tag, error) {
	var rows *sql.Rows
	var err error
	query := "select * from tag where title like '%" + title + "%'"
	stmt, err := prepare(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	return m.getModelList(rows)
}

// SaveTagAndMemo save tag and memo link
func (m tagRepository) SaveTagAndMemo(ctx context.Context, tagID int, memoID int) error {
	var err error
	query := "insert into tag_memo(tag_id, memo_id) values(?, ?)"
	stmt, err := prepare(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, tagID, memoID)
	return err
}

// GetAllByMemoID get all Tag Data By MemoID
func (m tagRepository) GetAllByMemoID(ctx context.Context, id int) ([]*model.Tag, error) {
	var rows *sql.Rows
	var err error
	query := "select tag.* from tag, tag_memo where tag_memo.memo_id = ? and tag.id = tag_memo.tag_id"
	stmt, err := prepare(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err = stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}

	return m.getModelList(rows)
}

// SearchMemoIDsByTitle search memo ids by tag's title
func (m tagRepository) SearchMemoIDsByTitle(ctx context.Context, title string) ([]int, error) {
	var rows *sql.Rows
	var err error
	query := "select tag_memo.memo_id as mid from tag, tag_memo where tag.id = tag_memo.tag_id and tag.title like '%" + title + "%'"
	stmt, err := prepare(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	list := []int{}
	for rows.Next() {
		var mid int64
		err := rows.Scan(&mid)
		if err != nil {
			return list, err
		}
		list = append(list, int(mid))
	}

	return list, nil
}

// getModelList get model list
func (m tagRepository) getModelList(rows *sql.Rows) ([]*model.Tag, error) {
	list := []*model.Tag{}
	for rows.Next() {
		tag := &model.Tag{}
		err := rows.Scan(&tag.ID, &tag.Title)
		if err != nil {
			return list, err
		}
		list = append(list, tag)
	}

	return list, nil
}
