package json

// Memo Memo's Out Entity
type Memo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Tag Tag's Out Entity
type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// PostMemoAndTagsResult Out Entity For PostMemoAndTags Result
type PostMemoAndTagsResult struct {
	Memo *Memo  `json:"memo"`
	Tags []*Tag `json:"tags"`
}

// SearchTagsAndMemosResult Out Entity For SearchTagsAndMemos Result
type SearchTagsAndMemosResult struct {
	Tags  []*Tag  `json:"tags"`
	Memos []*Memo `json:"memos"`
}
