package main

import (
	"net/http"
	"memo_sample/di"
	"memo_sample/infra"
)

func main() {
	infra.ConnectDB()
	defer infra.CloseDB()

	api := di.InjectDBAPI()
	http.HandleFunc("/", api.GetMemos)
	http.HandleFunc("/post", api.PostMemo)
	http.ListenAndServe(":8080", nil)
}
