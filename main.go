package main

import (
	"net/http"
	"memo_sample/di"
)

func main() {
	api := di.InjectAPI()
	http.HandleFunc("/", api.GetMemos)
	http.HandleFunc("/post", api.PostMemo)
	http.ListenAndServe(":8080", nil)
}