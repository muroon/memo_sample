package usecase

import (
	"testing"
	"context"
	"encoding/json"
	"memo_sample/adapter/memory"

)

func TestMemoPostAndGetSuccess(t *testing.T) {
	
	ctx := context.Background()

	repo := &memory.MemoRepository{}

	memo := NewMemo(repo)
	
	text := "First Memo"

	id, err := memo.Post(ctx, text)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess Post", err)
	}

	m, err := memo.Find(ctx, id)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess Get", err, id)
	}
	t.Logf("TestMemoPostAndGetSuccess Find MemoRepository id:%d, text:%s", m.ID, m.Text)

	list, err := memo.GetAll(ctx)	
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess GetAll", err)
	}
	for _, v := range list {
		t.Logf("TestMemoPostAndGetSuccess GetAll MemoRepository id:%d, text:%s", v.ID, v.Text)
	}
}

func TestMemoGetJsonSuccess(t *testing.T) {
	ctx := context.Background()

	repo := &memory.MemoRepository{}

	memo := NewMemo(repo)
	
	text := "Next Memo"

	id, err := memo.Post(ctx, text)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess Post", err)
	}

	m, err := memo.FindJSON(ctx, id)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess FindJSON", err, id)
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess Marshal", err)
		return
	}
	t.Logf("TestMemoPostAndGetSuccess Find MemoRepository json: %s", b)

	l, err := memo.GetAllJSON(ctx)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess GetAllJSON", err)
	}
	lb, err := json.Marshal(l)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess Marshal list", err)
		return
	}

	t.Logf("TestMemoPostAndGetSuccess GetAllJSON MemoRepository json: %s", lb)
}