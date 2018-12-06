package infra

import (
	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB DB接続
func ConnectDB() {
	err := (*GetDBM()).OpenDB("mysql", "root:@/memo_sample")
	if err != nil {
		panic(err)
	}
}

// CloseDB DB切断
func CloseDB() {
	(*GetDBM()).CloseDB()
}

// ConnectTestDB DB接続
func ConnectTestDB() {
	err := (*GetDBM()).OpenDB("mysql", "root:@/memo_sample_test")
	if err != nil {
		panic(err)
	}
}

// CloseTestDB DB切断
func CloseTestDB() {
	CloseDB()
}
