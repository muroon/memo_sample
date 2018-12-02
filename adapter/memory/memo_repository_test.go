package memory

import (
	"context"
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestMemoSaveInMemorySuccess(t *testing.T) {
	ctx := context.Background()

	repo := NewMemoRepository()

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
		t.Error("failed TestMemoSaveInMemorySuccess Save", err)
	}

	memoGet, err = repo.Get(ctx, memo.ID)
	if err != nil || memoGet.ID != memo.ID {
		t.Error("failed TestMemoSaveInMemorySuccess Get", err, memoGet.ID)
	}
	t.Logf("TestMemoSaveInMemorySuccess Get MemoRepository id:%d, text:%s", memoGet.ID, memoGet.Text)

	//　全件取得
	list, err := repo.GetAll(ctx)
	if err != nil || len(list) != 2 {
		t.Error("failed TestMemoSaveInMemorySuccess Get", err, len(list))
	}

	for _, v := range list {
		t.Logf("TestMemoSaveInMemorySuccess GetAll MemoRepository id:%d, text:%s", v.ID, v.Text)
	}
}
