package db

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func getTagRepositoryForTest() *TagRepository {
	return NewTagRepository(db)
}

func TestTagSaveInDBSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getTagRepositoryForTest()

	ctx := context.Background()

	// 1件名
	_, err := repo.Save(ctx, "Tag First")
	if err != nil {
		t.Error("failed TestTagSaveInTagrySuccess Save", err)
	}

}

func TestTagTransactionCommitSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getTagRepositoryForTest()

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

func TestTagAndMemoTransactionCommitSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()

	ctx := context.Background()

	ctx, err := repoM.Begin(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	_, err = repoM.Save(ctx, "Transaction Commit Multi Table Test(Memo)")
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	_, err = repoT.Save(ctx, "Transaction Commit Multi Table Test(Tag)")
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	_, err = repoM.Commit(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}
}
