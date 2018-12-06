package db

import (
	"memo_sample/infra"
)

// connectTestDB DB接続
func connectTestDB() {
	infra.ConnectTestDB()
}

// closeTestDB DB切断
func closeTestDB() {
	infra.CloseTestDB()
}
