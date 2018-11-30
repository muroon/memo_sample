package memory

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
	"context"
)

func TestMemoSaveInMemorySuccess(t *testing.T) {
	ctx := context.Background()

	repo := &MemoRepository{}

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
		t.Error("failed TestMemoSaveInMemorySuccess Save", err)
	}

	memoFind, err = repo.Find(ctx, memo.ID)
	if err != nil || memoFind.ID != memo.ID {
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