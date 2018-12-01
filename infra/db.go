package infra

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DbInfo DB情報
type DbInfo struct {
	DB *sql.DB
	Tx *sql.Tx
}

// NewDBInfo インフラ情報を渡す
func NewDBInfo() *DbInfo {
	return &DbInfo{DB: db, Tx: tx}
}

var db *sql.DB
var tx *sql.Tx

// ConnectDB DB接続
func ConnectDB() {
	dbconn, err := sql.Open("mysql", "root:@/memo_sample")
	if err != nil {
		panic(err)
	}
	db = dbconn
}

// CloseDB DB切断
func CloseDB() {
	db.Close()
}
