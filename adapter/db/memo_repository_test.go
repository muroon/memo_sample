package db

import (
	"context"
	"testing"
)

func getMemoRepositoryForTest() *MemoRepository {
	return NewMemoRepository(db)
}

func TestMemoSaveInDBSuccess(t *testing.T) {

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer closeTestDB()

	// 1件名
	memo, ctx, err := repo.Save(ctx, "First")
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess Save", err)
	}

	memoGet, ctx, err := repo.Get(ctx, memo.ID)
	if err != nil || memoGet.ID != memo.ID {
		t.Error("failed TestMemoSaveInMemorySuccess Get", err, memoGet.ID)
	}

	// 2件名
	memo, ctx, err = repo.Save(ctx, "Second")
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess GenerateID", err)
	}

	memoGet, ctx, err = repo.Get(ctx, memo.ID)
	if err != nil || memoGet.ID != memo.ID {
		t.Error("failed TestMemoSaveInMemorySuccess Get", err, memoGet.ID)
	}

	//　全件取得
	list, ctx, err := repo.GetAll(ctx)
	if err != nil {
		t.Error("failed TestMemoSaveInMemorySuccess Get", err, len(list))
	}

	for _, v := range list {
		t.Logf("TestMemoSaveInMemorySuccess GetAll MemoRepository id:%d, text:%s", v.ID, v.Text)
	}
}

func TestMemoTransactionCommitSuccess(t *testing.T) {

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer closeTestDB()

	ctx, err := repo.Begin(ctx)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	_, ctx, err = repo.Save(ctx, "Transaction Commit Test")
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

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer closeTestDB()

	ctx, err := repo.Begin(ctx)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	_, ctx, err = repo.Save(ctx, "Transaction Rollback Test")
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	repo.Rollback(ctx)
}

func TestMemoSearchSuccess(t *testing.T) {

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer closeTestDB()

	word := "Memo Search Test"
	_, ctx, err := repo.Save(ctx, word)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	word = "Memo"
	list, ctx, err := repo.Search(ctx, word)
	if err != nil {
		t.Error(err)
	}

	for _, m := range list {
		t.Log(m)
	}
}

func TestMemoGetAllByIDsSuccess(t *testing.T) {

	repo := getMemoRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer closeTestDB()

	word := "Dummy First"
	memo1, ctx, err := repo.Save(ctx, word)
	if err != nil {
		t.Error(err)
	}

	word = "Dummy Second"
	memo2, ctx, err := repo.Save(ctx, word)
	if err != nil {
		t.Error(err)
	}

	ids := []int{memo1.ID, memo2.ID}
	list, ctx, err := repo.GetAllByIDs(ctx, ids)
	if err != nil {
		t.Error(err)
	}

	for _, m := range list {
		t.Log(m)
	}
}
