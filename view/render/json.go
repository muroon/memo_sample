package render

import (
	"memo_sample/domain/model"
	"memo_sample/view/model/json"
)

// JSONRender api render interface
type JSONRender interface {
	ConvertError(err error, code int) *json.Error
	ConvertMemo(md *model.Memo) *json.Memo
	ConvertMemos(list []*model.Memo) []*json.Memo
	ConvertTag(md *model.Tag) *json.Tag
	ConvertTags(list []*model.Tag) []*json.Tag
	ConvertPostMemoAndTagsResult(memo *model.Memo, tags []*model.Tag) *json.PostMemoAndTagsResult
	ConvertSearchTagsAndMemosResult(memos []*model.Memo, tags []*model.Tag) *json.SearchTagsAndMemosResult
}
