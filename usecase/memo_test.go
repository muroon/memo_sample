package usecase

import (
	"testing"
	"context"
	"encoding/json"
	"memo_sample/adapter/memory"
	"memo_sample/usecase/input"
)

func TestMemoPostAndGetSuccess(t *testing.T) {
	
	ctx := context.Background()

	repo := &memory.MemoRepository{}

	memo := NewMemo(repo)
	
	text := "First Memo"

	ipt := &input.PostMemo{ Text: text }

	if err := memo.ValidatePost(*ipt); err != nil {
		t.Error("failed TestMemoPostAndGetSuccess ValidatePost", err)
	}

	id, err := memo.Post(ctx, *ipt)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess Post", err)
	}

	iptf := &input.FindMemo{ID: id}
	m, err := memo.Find(ctx, *iptf)
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

	ipt := &input.PostMemo{ Text: text }

	id, err := memo.Post(ctx, *ipt)
	if err != nil {
		t.Error("failed TestMemoPostAndGetSuccess Post", err)
	}

	iptf := &input.FindMemo{ID: id}
	m, err := memo.FindJSON(ctx, *iptf)
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