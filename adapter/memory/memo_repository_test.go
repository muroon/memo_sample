package memory

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
	"context"
	"memo_sample/domain/model"
)

func TestMemoFirstIDInMemorySuccess(t *testing.T) {
	ctx := context.Background()

	repo := &MemoRepository{}
	id, err := repo.GenerateID(ctx)
	if err != nil {
		t.Error("failed TestMemoFirstIDInMemorySuccess GenerateID")
	}

	t.Log("TestMemoFirstIDInMemorySuccess FirstID:", id)
}

func TestMemoSaveInMemorySuccess(t *testing.T) {
	ctx := context.Background()

	repo := &MemoRepository{}

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