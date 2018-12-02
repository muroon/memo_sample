package db

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

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
	return NewMemoRepository(db)
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

	memoGet, err := repo.Get(ctx, memo.ID)
	if err != nil || memoGet.ID != memo.ID {
		t.Error("failed TestMemoSaveInMemorySuccess Get", err, memoGet.ID)
	}
	t.Logf("TestMemoSaveInMemorySuccess Get MemoRepository id:%d, text:%s", memoGet.ID, memoGet.Text)

	// 2件名
	memo, err = repo.Save(ctx, "Second")
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess GenerateID", err)
	}

	memoGet, err = repo.Get(ctx, memo.ID)
	if err != nil || memoGet.ID != memo.ID {
		t.Error("failed TestMemoSaveInMemorySuccess Get", err, memoGet.ID)
	}
	t.Logf("TestMemoSaveInMemorySuccess Get MemoRepository id:%d, text:%s", memoGet.ID, memoGet.Text)

	//　全件取得
	list, err := repo.GetAll(ctx)
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess Get", err, len(list))
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

	_, err = repo.Save(ctx, "Transaction Commit Test")
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

func TestMemoSearchSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	word := "Memo Search Test"
	_, err := repo.Save(ctx, word)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	word = "Memo"
	list, err := repo.Search(ctx, word)
	if err != nil {
		t.Error(err)
	}

	for _, m := range list {
		t.Log(m)
	}
}

func TestMemoGetAllByIDsSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	word := "Dummy First"
	memo1, err := repo.Save(ctx, word)
	if err != nil {
		t.Error(err)
	}

	word = "Dummy Second"
	memo2, err := repo.Save(ctx, word)
	if err != nil {
		t.Error(err)
	}

	ids := []int{memo1.ID, memo2.ID}
	list, err := repo.GetAllByIDs(ctx, ids)
	if err != nil {
		t.Error(err)
	}

	for _, m := range list {
		t.Log(m)
	}
}
