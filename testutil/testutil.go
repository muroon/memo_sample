package testutil

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/memory"
	"memo_sample/domain/repository"
	"memo_sample/infra"
)

// NewTestManager test util
func NewTestManager() TestManager {
	return testManager{}
}

// TestManager test manager
type TestManager interface {
	GgetInMemoryRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository)
	GetDBRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository)
	ConnectTestDB() error
	CloseTestDB() error
}

// testManager test manager
type testManager struct {
}

func (t testManager) GgetInMemoryRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return memory.NewTransactionRepository(), memory.NewMemoRepository(), memory.NewTagRepository()
}

func (t testManager) GetDBRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return db.NewTransactionRepository(), db.NewMemoRepository(), db.NewTagRepository()
}

// connectTestDB DB接続
func (t testManager) ConnectTestDB() error {
	return (*infra.GetDBM()).ConnectTestDB()
}

// closeTestDB DB切断
func (t testManager) CloseTestDB() error {
	return (*infra.GetDBM()).CloseDB()
}
