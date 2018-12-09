package api

import (
	"context"
	"encoding/json"
	"fmt"
	"memo_sample/domain/model"
	"memo_sample/infra/logger"
	"memo_sample/usecase"
	"memo_sample/view/render"
	"net/http"
)

// NewPresenter new presenter
func NewPresenter(render render.JSONRender, log logger.Logger) usecase.Presenter {
	return presenter{render, log}
}

type presenter struct {
	render render.JSONRender
	log    logger.Logger
}

func (m presenter) ViewMemo(ctx context.Context, md *model.Memo) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	m.JSON(ctx, w, m.render.ConvertMemo(md))
}

func (m presenter) ViewMemoList(ctx context.Context, list []*model.Memo) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	m.JSON(ctx, w, m.render.ConvertMemos(list))
}

func (m presenter) ViewTag(ctx context.Context, md *model.Tag) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	m.JSON(ctx, w, m.render.ConvertTag(md))
}

func (m presenter) ViewTagList(ctx context.Context, list []*model.Tag) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	m.JSON(ctx, w, m.render.ConvertTags(list))
}

func (m presenter) ViewPostMemoAndTagsResult(ctx context.Context, memo *model.Memo, tags []*model.Tag) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	m.JSON(ctx, w, m.render.ConvertPostMemoAndTagsResult(memo, tags))
}

func (m presenter) ViewSearchTagsAndMemosResult(ctx context.Context, memos []*model.Memo, tags []*model.Tag) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	m.JSON(ctx, w, m.render.ConvertSearchTagsAndMemosResult(memos, tags))
}

func (m presenter) ViewError(ctx context.Context, err error) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	m.log.Errorf("%s", fmt.Sprintf("API: %T(%v)\n", err, err))

	m.JSON(ctx, w, m.render.ConvertError(err))
}

// JSON render json format
func (m presenter) JSON(ctx context.Context, w http.ResponseWriter, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
