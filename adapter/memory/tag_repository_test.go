package memory

import (
	"context"
	"fmt"
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestTagSaveInDBSuccess(t *testing.T) {

	repo := NewTagRepository()

	ctx := context.Background()

	me, err := repo.Save(ctx, "Tag First")
	if err != nil {
		t.Error("failed TestTagSaveInDBSuccess Save", err)
	}

	memo, err := repo.Get(ctx, me.ID)
	if err != nil {
		t.Error("failed TestTagSaveInDBSuccess Get", err)
	}

	t.Log(memo)
}

func TestTagAndMemoGetAllByMemoIDSuccess(t *testing.T) {

	repoTx := NewTransactionRepository()
	repoT := NewTagRepository()
	repoM := NewMemoRepository()

	ctx := context.Background()

	defer func() {
		if err := recover(); err != nil {
			repoTx.Rollback(ctx)
			t.Error(err)
		}
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

	t.Logf("TestTagAndMemoGetAllByMemoIDSuccess targetMemoID:%d", memo.ID)

	flag := false
	list, err := repoT.GetAllByMemoID(ctx, memo.ID)
	for _, tg := range list {
		if tg.ID == tag.ID {
			flag = true
			t.Log(tg)
		}
	}

	if !flag {
		panic(fmt.Errorf("GetAllByMemoID Error"))
	}
}

func TestTagAndMemoSearchMemoIDsByTitleSuccess(t *testing.T) {

	repoTx := NewTransactionRepository()
	repoT := NewTagRepository()
	repoM := NewMemoRepository()

	ctx := context.Background()

	defer func() {
		if err := recover(); err != nil {
			repoTx.Rollback(ctx)
			t.Error(err)
		}
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

	ctx, err = repoTx.Commit(ctx)
	if err != nil {
		panic(err)
	}

	flag := false
	list, err := repoT.SearchMemoIDsByTitle(ctx, tag.Title)
	for _, id := range list {
		if id == memo.ID {
			flag = true
		}
	}

	if !flag {
		t.Error(fmt.Errorf("SearchMemoIDsByTitle Error"))
	}
}
