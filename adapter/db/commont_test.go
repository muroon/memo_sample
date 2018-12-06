package db

import (
	"memo_sample/infra"
)

// ConnectDB DB接続
func connectTestDB() {
	infra.ConnectTestDB()
}

// CloseDB DB切断
func closeTestDB() {
	infra.CloseTestDB()
}
