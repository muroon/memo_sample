package api

import (
	"context"
	"encoding/json"
	"memo_sample/usecase"
	"memo_sample/usecase/input"
	"net/http"
)

// NewAPI Get API instance
func NewAPI(
	memo usecase.Memo,
) API {
	return API{
		memo,
	}
}

// API api instance
type API struct {
	memo usecase.Memo
}

// PostMemo post new memo
func (a API) PostMemo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ipt := &input.PostMemo{Text: r.URL.Query().Get("text")}
	id, err := a.memo.Post(ctx, *ipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	iptf := &input.GetMemo{ID: id}
	v, err := a.memo.GetMemo(ctx, *iptf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	a.JSON(ctx, w, v)
}

// GetMemos get all memo
func (a API) GetMemos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	v, err := a.memo.GetAllMemoList(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	a.JSON(ctx, w, v)
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
