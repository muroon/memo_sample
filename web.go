package main

import (
	"flag"
	loggersub "memo_sample/adapter/logger"
	"memo_sample/di"
	"memo_sample/infra/database"
	"net"
	"net/http"
)

func main() {
	ping := flag.Bool("ping", false, "check ping")
	flag.Parse()

	defer func() {
		_ = (*database.GetDBM()).CloseDB()
	}()

	err := (*database.GetDBM()).ConnectDB()
	if err != nil {
		loggersub.NewLogger().Errorf("db open error: %#+v\n", err)
		return
	}

	interceptor := func(h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request)	{
			var err error
			if *ping {
				err = (*database.GetDBM()).PingDB()
			}
			if err != nil {
				loggersub.NewLogger().Errorf("db open error: %#+v\n", err)
				panic(err)
			}
			h(w, r)
		}
	}

	loggersub.NewLogger().Debugf("main called. ping check:%v\n", *ping)

	api := di.InjectAPIServer()
	http.HandleFunc("/", interceptor(api.GetMemos))
	http.HandleFunc("/post", interceptor(api.PostMemo))
	http.HandleFunc("/post/memo_tags", interceptor(api.PostMemoAndTags))
	http.HandleFunc("/search/tags_memos", interceptor(api.SearchTagsAndMemos))
	lin, err := net.Listen("tcp4", ":8080")
	if err != nil {
		loggersub.NewLogger().Errorf("Listen error: %#+v\n", err)
	}
	defer lin.Close()

	s := new(http.Server)
	s.Serve(lin)
}
