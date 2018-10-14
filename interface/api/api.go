package api

import (
	"context"
	"encoding/json"
	"net/http"
	"memo_sample/usecase"
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
	memo        usecase.Memo
}

// PostMemo post new memo
func (a API) PostMemo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := a.memo.Post(ctx, r.URL.Query().Get("text"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	v, err := a.memo.FindJSON(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	a.JSON(ctx, w, v)
}

// GetMemos get all memo
func (a API) GetMemos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	v, err := a.memo.GetAllJSON(ctx)
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