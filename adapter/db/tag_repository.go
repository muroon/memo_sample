package db

import (
	"context"
	"database/sql"
	"memo_sample/domain/model"
	"memo_sample/domain/repository"

	"net/http"
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
	defer stmt.Close()
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	res, err = stmt.ExecContext(ctx, title)
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	return m.Get(ctx, int(id))
}

// Get get Tag Data by ID
func (m tagRepository) Get(ctx context.Context, id int) (*model.Tag, error) {
	tag := new(model.Tag)
	var err error
	query := "select * from tag where id = ?"
	stmt, err := prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&tag.ID, &tag.Title)
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	return tag, err
}

// GetAll get all Tag Data
func (m tagRepository) GetAll(ctx context.Context) ([]*model.Tag, error) {
	var rows *sql.Rows
	var err error
	query := "select * from tag"
	stmt, err := prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	return m.getModelList(rows)
}

// Search search list by title
func (m tagRepository) Search(ctx context.Context, title string) ([]*model.Tag, error) {
	var rows *sql.Rows
	var err error
	query := "select * from tag where title like '%" + title + "%'"
	stmt, err := prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	return m.getModelList(rows)
}

// SaveTagAndMemo save tag and memo link
func (m tagRepository) SaveTagAndMemo(ctx context.Context, tagID int, memoID int) error {
	var err error
	query := "insert into tag_memo(tag_id, memo_id) values(?, ?)"
	stmt, err := prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return errm.Wrap(err, http.StatusInternalServerError)
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
	defer stmt.Close()
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	rows, err = stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	return m.getModelList(rows)
}

// SearchMemoIDsByTitle search memo ids by tag's title
func (m tagRepository) SearchMemoIDsByTitle(ctx context.Context, title string) ([]int, error) {
	var rows *sql.Rows
	var err error
	query := "select tag_memo.memo_id as mid from tag, tag_memo where tag.id = tag_memo.tag_id and tag.title like '%" + title + "%'"
	stmt, err := prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return nil, errm.Wrap(err, http.StatusInternalServerError)
	}

	list := []int{}
	for rows.Next() {
		var mid int64
		err := rows.Scan(&mid)
		if err != nil {
			return list, errm.Wrap(err, http.StatusInternalServerError)
		}
		list = append(list, int(mid))
	}

	return list, nil
}

// getModelList get model list
func (m tagRepository) getModelList(rows *sql.Rows) ([]*model.Tag, error) {
	list := make([]*model.Tag, 0)
	for rows.Next() {
		tag := &model.Tag{}
		err := rows.Scan(&tag.ID, &tag.Title)
		if err != nil {
			return list, errm.Wrap(err, http.StatusInternalServerError)
		}
		list = append(list, tag)
	}

	return list, nil
}
