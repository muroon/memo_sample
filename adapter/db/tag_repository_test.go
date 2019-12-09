package db

import (
	"context"
	"fmt"
	"memo_sample/domain/repository"
	"testing"
)

func getTagRepositoryForTest() repository.TagRepository {
	return NewTagRepository()
}

func TestTagSaveInDBSuccess(t *testing.T) {

	repo := getTagRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer closeTestDB()

	tag, err := repo.Save(ctx, "Tag First")
	if err != nil {
		t.Error("failed TestTagSaveInTagrySuccess Save", err)
	}

	tagGet, err := repo.Get(ctx, tag.ID)
	if err != nil || tagGet.ID != tag.ID {
		t.Error("failed TestTagSaveInDBSuccess Get", err, tag.ID)
	}
}

func TestTagTransactionCommitSuccess(t *testing.T) {
	repo := getTagRepositoryForTest()
	repoTx := getTransactionRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer func() {
		if err := recover(); err != nil {
			_, _ = repoTx.Rollback(ctx)
			t.Error(err)
		}
		closeTestDB()
	}()

	ctx, err := repoTx.Begin(ctx)
	if err != nil {
		panic(err)
	}

	_, err = repo.Save(ctx, "Transaction Commit Test")
	if err != nil {
		panic(err)
	}

	_, err = repoTx.Commit(ctx)
	if err != nil {
		panic(err)
	}
}

func TestTagAndMemoTransactionCommitSuccess(t *testing.T) {

	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()
	repoTx := getTransactionRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer func() {
		if err := recover(); err != nil {
			_, _ = repoTx.Rollback(ctx)
			t.Error(err)
		}
		closeTestDB()
	}()

	ctx, err := repoTx.Begin(ctx)
	if err != nil {
		panic(err)
	}

	memo, err := repoM.Save(ctx, "Transaction Commit Memo")
	if err != nil {
		panic(err)
	}

	tag, err := repoT.Save(ctx, "Transaction Commit Tag")
	if err != nil {
		panic(err)
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		panic(err)
	}

	_, err = repoTx.Commit(ctx)
	if err != nil {
		panic(err)
	}
}

func TestTagAndMemoTransactionRollbackSuccess(t *testing.T) {

	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()
	repoTx := getTransactionRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer func() {
		if err := recover(); err != nil {
			_, _ = repoTx.Rollback(ctx)
			t.Error(err)
		}
		closeTestDB()
	}()

	ctx, err := repoTx.Begin(ctx)
	if err != nil {
		panic(err)
	}

	memo, err := repoM.Save(ctx, "Transaction Rollback Memo")
	if err != nil {
		panic(err)
	}

	tag, err := repoT.Save(ctx, "Transaction Rollback Tag")
	if err != nil {
		panic(err)
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		panic(err)
	}

	// 強制的にロールバック
	_, _ = repoTx.Rollback(ctx)
}

func TestTagAndMemoGetAllByMemoIDSuccess(t *testing.T) {
	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()
	repoTx := getTransactionRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer func() {
		if err := recover(); err != nil {
			_, _ = repoTx.Rollback(ctx)
			t.Error(err)
		}
		closeTestDB()
	}()

	ctx, err := repoTx.Begin(ctx)
	if err != nil {
		panic(err)
	}

	memo, err := repoM.Save(ctx, "GetAllByMemoID Test Memo")
	if err != nil {
		panic(err)
	}

	tag, err := repoT.Save(ctx, "GetAllByMemoID Test Tag")
	if err != nil {
		panic(err)
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		panic(err)
	}

	ctx, err = repoTx.Commit(ctx)
	if err != nil {
		panic(err)
	}

	flag := false
	list, err := repoT.GetAllByMemoID(ctx, memo.ID)
	if err !=nil {
		panic(err)
	}
	for _, tg := range list {
		if tg.ID == tag.ID {
			flag = true
		}
	}

	if !flag {
		panic(fmt.Errorf("GetAllByMemoID Error"))
	}
}

func TestTagAndMemoSearchMemoIDsByTitleSuccess(t *testing.T) {

	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()
	repoTx := getTransactionRepositoryForTest()

	ctx := context.Background()

	connectTestDB()
	defer func() {
		if err := recover(); err != nil {
			_, _ = repoTx.Rollback(ctx)
			t.Error(err)
		}
		closeTestDB()
	}()

	ctx, err := repoTx.Begin(ctx)
	if err != nil {
		panic(err)
	}

	memo, err := repoM.Save(ctx, "SearchMemoIDsByTitle Test Memo")
	if err != nil {
		panic(err)
	}

	tag, err := repoT.Save(ctx, "SearchMemoIDsByTitle Test Tag")
	if err != nil {
		panic(err)
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		panic(err)
	}

	tag2, err := repoT.Get(ctx, tag.ID)
	if err != nil {
		panic(err)
	}
	t.Log(tag2)

	ctx, err = repoTx.Commit(ctx)
	if err != nil {
		panic(err)
	}

	flag := false
	list, err := repoT.SearchMemoIDsByTitle(ctx, tag.Title)
	if err != nil {
		panic(err)
	}
	for _, id := range list {
		if id == memo.ID {
			flag = true
		}
	}

	if !flag {
		panic(fmt.Errorf("SearchMemoIDsByTitle Error"))
	}
}
