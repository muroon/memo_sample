package view

import (
	"fmt"
	"memo_sample/domain/model"
	"memo_sample/interface/api"
	"memo_sample/view/model/json"
)

// NewAPIRender
func NewAPIRender() api.APIRender {
	return apiRender{}
}

type apiRender struct {
}

func (m apiRender) ConvertMemoJSON(md *model.Memo) *json.Memo {
	mj := &json.Memo{
		ID:   md.ID,
		Text: md.Text,
	}
	return mj
}

func (m apiRender) ConvertMemoJSONList(list []*model.Memo) []*json.Memo {
	listJSON := []*json.Memo{}
	for _, v := range list {
		listJSON = append(listJSON, m.ConvertMemoJSON(v))
	}
	return listJSON
}

func (m apiRender) ConvertTagJSON(md *model.Tag) *json.Tag {
	mj := &json.Tag{
		ID:    md.ID,
		Title: md.Title,
	}
	return mj
}

func (m apiRender) ConvertTagJSONList(list []*model.Tag) []*json.Tag {
	listJSON := []*json.Tag{}
	for _, v := range list {
		listJSON = append(listJSON, m.ConvertTagJSON(v))
	}
	return listJSON
}

func (m apiRender) ConvertPostMemoAndTagsResultJSON(memo *model.Memo, tags []*model.Tag) *json.PostMemoAndTagsResult {

	return &json.PostMemoAndTagsResult{
		Memo: m.ConvertMemoJSON(memo),
		Tags: m.ConvertTagJSONList(tags),
	}
}

func (m apiRender) ConvertSearchTagsAndMemosResultJSON(memos []*model.Memo, tags []*model.Tag) *json.SearchTagsAndMemosResult {

	return &json.SearchTagsAndMemosResult{
		Tags:  m.ConvertTagJSONList(tags),
		Memos: m.ConvertMemoJSONList(memos),
	}
}

func (m apiRender) ConvertError(err error) *json.Error {
	mess := fmt.Sprintf("API: %T(%v)\n", err, err)

	return &json.Error{mess}
}
