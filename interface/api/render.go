package api

import (
	"memo_sample/domain/model"
	"memo_sample/view/model/json"
)

// APIRender
type APIRender interface {
	ConvertError(err error) *json.Error
	ConvertMemoJSON(md *model.Memo) *json.Memo
	ConvertMemoJSONList(list []*model.Memo) []*json.Memo
	ConvertTagJSON(md *model.Tag) *json.Tag
	ConvertTagJSONList(list []*model.Tag) []*json.Tag
	ConvertPostMemoAndTagsResultJSON(memo *model.Memo, tags []*model.Tag) *json.PostMemoAndTagsResult
	ConvertSearchTagsAndMemosResultJSON(memos []*model.Memo, tags []*model.Tag) *json.SearchTagsAndMemosResult
}
