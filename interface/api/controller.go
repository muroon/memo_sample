package api

import (
	"memo_sample/infra/logger"
	"memo_sample/usecase"
	"memo_sample/usecase/input"
	"net/http"
)

// NewAPI Get API instance
func NewAPI(it usecase.Interactor, log logger.Logger) API {
	return controller{it, log}
}

// API api instance
type API interface {
	PostMemo(w http.ResponseWriter, r *http.Request)
	GetMemos(w http.ResponseWriter, r *http.Request)
	PostMemoAndTags(w http.ResponseWriter, r *http.Request)
	SearchTagsAndMemos(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	it  usecase.Interactor
	log logger.Logger
}

// PostMemo post new memo
func (c controller) PostMemo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = setResponseWriter(ctx, w)

	ipt := &input.PostMemo{Text: r.URL.Query().Get("text")}
	c.it.PostMemo(ctx, *ipt)
}

// GetMemos get all memo
func (c controller) GetMemos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = setResponseWriter(ctx, w)

	c.it.GetMemos(ctx)
}

// PostMemoAndTags save memo and tags
func (c controller) PostMemoAndTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	r.ParseForm()
	text := r.FormValue("memo_text")
	titles := r.Form["tag_titles[]"]

	ctx = setResponseWriter(ctx, w)

	ipt := &input.PostMemoAndTags{
		MemoText:  text,
		TagTitles: titles,
	}

	c.it.PostMemoAndTags(ctx, *ipt)
}

// SearchTagsAndMemos save memo and tags
func (c controller) SearchTagsAndMemos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = setResponseWriter(ctx, w)

	title := r.URL.Query().Get("tag_title")

	ipt := &input.SearchTagsAndMemos{TagTitle: title}
	c.it.SearchTagsAndMemos(ctx, *ipt)
}
