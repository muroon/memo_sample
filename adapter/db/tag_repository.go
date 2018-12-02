package db

import (
	"context"
	"database/sql"
)

// NewTagRepository get repository
func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{}
}

// TagRepository Tag's Repository Sub
type TagRepository struct{}

// Begin begin transaction
func (m *TagRepository) Begin(ctx context.Context) (context.Context, error) {
	return begin(ctx)
}

// Rollback rollback transaction
func (m *TagRepository) Rollback(ctx context.Context) (context.Context, error) {
	return rollback(ctx)
}

// Commit commit transaction
func (m *TagRepository) Commit(ctx context.Context) (context.Context, error) {
	return commit(ctx)
}

// Save save Tag Data
//func (m *TagRepository) Save(ctx context.Context, title string) (*model.Tag, error) {
func (m *TagRepository) Save(ctx context.Context, title string) (int, error) {
	var err error
	var res sql.Result
	if isTx(ctx) {
		res, err = tx.Exec("insert into tag(title) values(?)", title)

	} else {
		res, err = db.Exec("insert into tag(title) values(?)", title)
	}

	id, err := res.LastInsertId()
	if err != nil {
		//return nil, err
		return 0, err
	}

	//return m.Get(ctx, int(id))
	return int(id), err
}
