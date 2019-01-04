package view

import (
	"fmt"
	"memo_sample/domain/model"
	"memo_sample/view/model/json"
	"memo_sample/view/render"
)

// NewJSONRender new json render
func NewJSONRender() render.JSONRender {
	return jsonRender{}
}

type jsonRender struct {
}

func (m jsonRender) ConvertMemo(md *model.Memo) *json.Memo {
	mj := &json.Memo{
		ID:   md.ID,
		Text: md.Text,
	}
	return mj
}

func (m jsonRender) ConvertMemos(list []*model.Memo) []*json.Memo {
	listJSON := []*json.Memo{}
	for _, v := range list {
		listJSON = append(listJSON, m.ConvertMemo(v))
	}
	return listJSON
}

func (m jsonRender) ConvertTag(md *model.Tag) *json.Tag {
	mj := &json.Tag{
		ID:    md.ID,
		Title: md.Title,
	}
	return mj
}

func (m jsonRender) ConvertTags(list []*model.Tag) []*json.Tag {
	listJSON := []*json.Tag{}
	for _, v := range list {
		listJSON = append(listJSON, m.ConvertTag(v))
	}
	return listJSON
}

func (m jsonRender) ConvertPostMemoAndTagsResult(memo *model.Memo, tags []*model.Tag) *json.PostMemoAndTagsResult {

	return &json.PostMemoAndTagsResult{
		Memo: m.ConvertMemo(memo),
		Tags: m.ConvertTags(tags),
	}
}

func (m jsonRender) ConvertSearchTagsAndMemosResult(memos []*model.Memo, tags []*model.Tag) *json.SearchTagsAndMemosResult {

	return &json.SearchTagsAndMemosResult{
		Tags:  m.ConvertTags(tags),
		Memos: m.ConvertMemos(memos),
	}
}

func (m jsonRender) ConvertError(err error, code int) *json.Error {
	mess := fmt.Sprintf("API: %T(%v)\n", err, err)

	return &json.Error{Code: code, Msg: mess}
}
