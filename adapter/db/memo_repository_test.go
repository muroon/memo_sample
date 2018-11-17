package db

import (
	"testing"
	"context"
	"memo_sample/domain/model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB;

// ConnectDB DB接続
func connectTestDB() {
	dbconn, err := sql.Open("mysql", "root:@/memo_sample_test")
	if err != nil {
		panic(err)
	}
	db = dbconn
}

// CloseDB DB切断
func closeTestDB() {
	db.Close() 
}

func getMemoRepositoryForTest() *MemoRepository {
	return &MemoRepository{DB: db}
}

func TestMemoFirstIDInDBSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	id, err := repo.GenerateID(ctx)
	if err != nil {
		t.Error("failed TestMemoFirstIDInMemorySuccess GenerateID")
	}

	t.Log("TestMemoFirstIDInMemorySuccess FirstID:", id)
}

func TestMemoSaveInDBSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	// 1件名
	id, err := repo.GenerateID(ctx)
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess GenerateID", err)
	}

	memo := &model.Memo{
		ID: id,
		Text: "First",
	}
	repo.Save(ctx, memo)

	memoFind, err := repo.Find(ctx, id)
	if err != nil || memoFind.ID != id {
		t.Error("failed TestMemoSaveInMemorySuccess Find", err, memoFind.ID)
	}
	t.Logf("TestMemoSaveInMemorySuccess Find MemoRepository id:%d, text:%s", memoFind.ID, memoFind.Text)

	// 2件名
	id, err = repo.GenerateID(ctx)
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess GenerateID", err)
	}

	memo = &model.Memo{
		ID: id,
		Text: "Second",
	}
	repo.Save(ctx, memo)

	memoFind, err = repo.Find(ctx, id)
	if err != nil || memoFind.ID != id {
		t.Error("failed TestMemoSaveInMemorySuccess Find", err, memoFind.ID)
	}
	t.Logf("TestMemoSaveInMemorySuccess Find MemoRepository id:%d, text:%s", memoFind.ID, memoFind.Text)

	//　全件取得
	list, err := repo.GetAll(ctx)
	if err != nil || len(list) != 2 {
		t.Error("failed TestMemoSaveInMemorySuccess Find", err, len(list))
	}

	for _, v := range list {
		t.Logf("TestMemoSaveInMemorySuccess GetAll MemoRepository id:%d, text:%s", v.ID, v.Text)
	}
}