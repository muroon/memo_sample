package db

import (
	"testing"
	"context"
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

func TestMemoSaveInDBSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	// 1件名
	memo, err := repo.Save(ctx, "First")
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess Save", err)
	}

	memoFind, err := repo.Find(ctx, memo.ID)
	if err != nil || memoFind.ID != memo.ID {
		t.Error("failed TestMemoSaveInMemorySuccess Find", err, memoFind.ID)
	}
	t.Logf("TestMemoSaveInMemorySuccess Find MemoRepository id:%d, text:%s", memoFind.ID, memoFind.Text)

	// 2件名
	memo, err = repo.Save(ctx, "Second")
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess GenerateID", err)
	}

	memoFind, err = repo.Find(ctx, memo.ID)
	if err != nil || memoFind.ID != memo.ID {
		t.Error("failed TestMemoSaveInMemorySuccess Find", err, memoFind.ID)
	}
	t.Logf("TestMemoSaveInMemorySuccess Find MemoRepository id:%d, text:%s", memoFind.ID, memoFind.Text)

	//　全件取得
	list, err := repo.GetAll(ctx)
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess Find", err, len(list))
	}

	for _, v := range list {
		t.Logf("TestMemoSaveInMemorySuccess GetAll MemoRepository id:%d, text:%s", v.ID, v.Text)
	}
}

func TestMemoTransactionCommitSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	ctx, err := repo.Begin(ctx)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	_, err = repo.Save(ctx, "Transaction Commit Test!")
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	_, err = repo.Commit(ctx)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}
}

func TestMemoTransactionRollbackSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	ctx, err := repo.Begin(ctx)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	_, err = repo.Save(ctx, "Transaction Rollback Test")
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	repo.Rollback(ctx)
}