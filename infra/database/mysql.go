package database

import (
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB DB接続
func (m *dbm) ConnectDB() error {
	return m.openDB("mysql", "root:@/memo_sample")
}

// ConnectTestDB Test用 DB接続
func (m *dbm) ConnectTestDB() error {
	return m.openDB("mysql", "root:@/memo_sample_test")
}

// CloseDB DB切断
func (m *dbm) CloseDB() error {
	return m.closeDB()
}
