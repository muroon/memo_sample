package db

import (
	"memo_sample/domain/repository"
)

// connectTestDB DB接続
func connectTestDB() {
	if err := (*dbm).ConnectTestDB(); err != nil {
		panic(err)
	}
}

// closeTestDB DB切断
func closeTestDB() {
	_ = (*dbm).CloseDB()
}

// getTransactionRepositoryForTest get TransactionRepository
func getTransactionRepositoryForTest() repository.TransactionRepository {
	return NewTransactionRepository()
}
