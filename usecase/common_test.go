package usecase

import (
	"memo_sample/domain/repository"
	"memo_sample/testutil"
)

var testManager testutil.TestManager

func init() {
	testManager = testutil.NewTestManager()
}

// getInMemoryRepository get memory repository
func getInMemoryRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return testManager.GgetInMemoryRepository()
}

// getDBRepository get db repository
func getDBRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return testManager.GetDBRepository()
}

// connectTestDB DB接続
func connectTestDB() {
	testManager.ConnectTestDB()
}

// closeTestDB DB切断
func closeTestDB() {
	testManager.CloseTestDB()
}
