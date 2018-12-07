package usecase

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/memory"
	"memo_sample/domain/repository"
	"memo_sample/infra"
)

func getInMemoryRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return memory.NewTransactionRepository(), memory.NewMemoRepository(), memory.NewTagRepository()
}

func getDBRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return db.NewTransactionRepository(), db.NewMemoRepository(), db.NewTagRepository()
}

// connectTestDB DB接続
func connectTestDB() {
	infra.ConnectTestDB()
}

// closeTestDB DB切断
func closeTestDB() {
	infra.CloseTestDB()
}
