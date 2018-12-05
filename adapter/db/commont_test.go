package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB DB接続
func connectTestDB() {
	dbconn, err := sql.Open("mysql", "root:@/memo_sample_test")
	if err != nil {
		panic(err)
	}
	db = dbconn
}

// CloseDB DB切断
func closeTestDB() {
	stmt.Close()
	db.Close()
}
