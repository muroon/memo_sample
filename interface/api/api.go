package api

import (
	"context"
	"encoding/json"
	"fmt"
	"memo_sample/infra"
	"memo_sample/usecase"
	"memo_sample/usecase/input"
	"net/http"
)

// NewAPI Get API instance
func NewAPI(memo usecase.Memo, render APIRender, log infra.Log) API {
	return API{memo, render, log}
}

// API api instance
type API struct {
	memo   usecase.Memo
	render APIRender
	log    infra.Log
}

// PostMemo post new memo
func (a API) PostMemo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ipt := &input.PostMemo{Text: r.URL.Query().Get("text")}
	id, err := a.memo.Post(ctx, *ipt)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}

	iptf := &input.GetMemo{ID: id}
	memo, err := a.memo.GetMemo(ctx, *iptf)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}

	a.JSON(ctx, w, a.render.ConvertMemoJSON(memo))
}

// GetMemos get all memo
func (a API) GetMemos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	memos, err := a.memo.GetAllMemoList(ctx)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}

	a.JSON(ctx, w, a.render.ConvertMemoJSONList(memos))
}

// PostMemoAndTags save memo and tags
func (a API) PostMemoAndTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	r.ParseForm()
	text := r.FormValue("memo_text")
	titles := r.Form["tag_titles[]"]

	ipt := &input.PostMemoAndTags{
		MemoText:  text,
		TagTitles: titles,
	}
	memo, tags, err := a.memo.PostMemoAndTags(ctx, *ipt)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}

	a.JSON(ctx, w, a.render.ConvertPostMemoAndTagsResultList(memo, tags))
}

// SearchTagsAndMemos save memo and tags
func (a API) SearchTagsAndMemos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	title := r.URL.Query().Get("tag_title")

	ipt := &input.SearchTagsAndMemos{TagTitle: title}
	memos, tags, err := a.memo.SearchTagsAndMemos(ctx, *ipt)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}

	a.JSON(ctx, w, a.render.ConvertSearchTagsAndMemosResultJSONList(memos, tags))
}

// HandleError handle error
func (a API) HandleError(ctx context.Context, w http.ResponseWriter, err error) {

	a.log.Errorf("%s", fmt.Sprintf("API: %T(%v)\n", err, err))

	a.JSON(ctx, w, a.render.ConvertError(err))
}

// JSON render json format
func (a API) JSON(ctx context.Context, w http.ResponseWriter, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
