package usecase

import (
	"context"
	"encoding/json"
	"memo_sample/usecase/input"
	"strings"
	"testing"
)

func TestMemoPostAndGetMemoInMemorySuccess(t *testing.T) {
	ctx := context.Background()

	// InMemoryでテストした場合
	memo := NewMemo(getInMemoryRepository())

	text := "Next Memo"

	ipt := &input.PostMemo{Text: text}

	id, err := memo.Post(ctx, *ipt)
	if err != nil {
		t.Error(err)
	}

	iptf := &input.GetMemo{ID: id}
	m, err := memo.GetMemo(ctx, *iptf)
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("TestMemoPostAndGetSuccess Get MemoRepository json: %s", b)

	l, err := memo.GetAllMemoList(ctx)
	if err != nil {
		t.Error(err)
	}
	lb, err := json.Marshal(l)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("TestMemoPostAndGetSuccess GetAllJSON MemoRepository json: %s", lb)
}

func TestMemoPostAndGetMemoInDBSuccess(t *testing.T) {
	ctx := context.Background()

	// DBでテストした場合
	memo := NewMemo(getDBRepository())

	connectTestDB()
	defer closeTestDB()

	text := "Next Memo"

	ipt := &input.PostMemo{Text: text}

	id, err := memo.Post(ctx, *ipt)
	if err != nil {
		t.Error(err)
	}

	iptf := &input.GetMemo{ID: id}
	m, err := memo.GetMemo(ctx, *iptf)
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("TestMemoPostAndGetSuccess Get MemoRepository json: %s", b)

	l, err := memo.GetAllMemoList(ctx)
	if err != nil {
		t.Error(err)
	}
	lb, err := json.Marshal(l)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("TestMemoPostAndGetSuccess GetAllJSON MemoRepository json: %s", lb)
}

func TestMemoPostMemoAndTagSuccess(t *testing.T) {
	ctx := context.Background()

	memo := NewMemo(getInMemoryRepository())

	memoText := "Post Memo And Tag Test"
	tagTitles := []string{"MemoTest", "UnitTest"}
	ipt := &input.PostMemoAndTags{MemoText: memoText, TagTitles: tagTitles}

	res, err := memo.PostMemoAndTags(ctx, *ipt)
	if err != nil {
		t.Error(err)
	}

	if res.Memo.Text != memoText {
		t.Errorf("Memo Save Error: %s", res.Memo.Text)
	}

	ok := []int{}
	for _, tg := range res.Tags {
		for _, title := range tagTitles {
			if tg.Title == title {
				ok = append(ok, 1)
			}
		}
	}

	if len(ok) != 2 {
		t.Errorf("Tag Save Error: %s", res.Memo.Text)
	}

	// 検索して取得して確認
	ipt2 := &input.GetTagsByMemo{ID: res.Memo.ID}
	tgs2, err := memo.GetTagsByMemo(ctx, *ipt2)

	ok = []int{}
	for _, tg := range tgs2 {
		for _, title := range tagTitles {
			if tg.Title == title {
				ok = append(ok, 1)
			}
		}
	}

	if len(ok) < 2 {
		t.Errorf("Tag Save Error: %s", res.Memo.Text)
	}
}

func TestMemoSearchTagsAndMemosSuccess(t *testing.T) {
	ctx := context.Background()

	memo := NewMemo(getInMemoryRepository())

	// test deta post
	memoTexts := []string{"SearchTagsAndMemos 1", "SearchTagsAndMemos 2"}
	tagTitle := "SearchTagsAndMemos"
	tagTitles := []string{tagTitle}
	for _, memoText := range memoTexts {
		ipt1 := &input.PostMemoAndTags{MemoText: memoText, TagTitles: tagTitles}
		_, err := memo.PostMemoAndTags(ctx, *ipt1)
		if err != nil {
			t.Error(err)
		}
	}

	ipt := &input.SearchTagsAndMemos{TagTitle: tagTitle}

	res, err := memo.SearchTagsAndMemos(ctx, *ipt)
	if err != nil {
		t.Error(err)
	}

	// check Tag
	for _, tag := range res.Tags {
		if strings.Index(tag.Title, tagTitle) == -1 {
			t.Errorf("Tag And Memo Save Error. tag.Title:%s", tag.Title)
		}
	}

	// check Memo
	ok := []int{}
	for _, mm := range res.Memos {
		for _, memoText := range memoTexts {
			if mm.Text == memoText {
				ok = append(ok, 1)
			}
		}
	}

	if len(ok) < 2 {
		t.Error("Tag And Memo Save Error")
	}
}
