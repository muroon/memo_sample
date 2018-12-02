package input

// PostMemo Input Entity For Posting Memo
type PostMemo struct {
	Text string
}

// PostMemoAndTags Input Entity For Posting Memo And Title
type PostMemoAndTags struct {
	MemoText  string
	TagTitles []string
}

// GetMemo Input Entity For Get Memo
type GetMemo struct {
	ID int
}

// GetTagsByMemo Input Entity For GetTagsByMemo
type GetTagsByMemo struct {
	ID int
}

// SearchTagsAndMemos Input Entity For SearchTagsAndMemos
type SearchTagsAndMemos struct {
	TagTitle string
}
