package main

import (
	loggersub "memo_sample/adapter/logger"
	"memo_sample/di"
	"memo_sample/infra/database"
	"net/http"
)

func main() {
	defer func() {
		_ = (*database.GetDBM()).CloseDB()
	}()

	err := (*database.GetDBM()).ConnectDB()
	if err != nil {
		loggersub.NewLogger().Errorf("db open error: %#+v\n", err)
		return
	}

	api := di.InjectAPIServer()
	http.HandleFunc("/", api.GetMemos)
	http.HandleFunc("/post", api.PostMemo)
	http.HandleFunc("/post/memo_tags", api.PostMemoAndTags)
	http.HandleFunc("/search/tags_memos", api.SearchTagsAndMemos)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		loggersub.NewLogger().Errorf("ListenAndServe error: %#+v\n", err)
	}
}
