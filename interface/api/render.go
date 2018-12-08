package api

import (
	"memo_sample/domain/model"
	"memo_sample/view/model/json"
)

type APIRender interface {
	ConvertMemoJSON(md *model.Memo) *json.Memo
	ConvertMemoJSONList(list []*model.Memo) []*json.Memo
	ConvertTagJSON(md *model.Tag) *json.Tag
	ConvertTagJSONList(list []*model.Tag) []*json.Tag
	ConvertPostMemoAndTagsResultList(memo *model.Memo, tags []*model.Tag) *json.PostMemoAndTagsResult
	ConvertSearchTagsAndMemosResultJSONList(memos []*model.Memo, tags []*model.Tag) *json.SearchTagsAndMemosResult
	ConvertError(err error) *json.Error
}
