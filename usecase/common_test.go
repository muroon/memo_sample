package usecase

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/memory"
	"memo_sample/domain/repository"
	"memo_sample/infra"
)

func getInMemoryRepository() (repository.MemoRepository, repository.TagRepository) {
	return memory.NewMemoRepository(), memory.NewTagRepository()
}

func getDBRepository() (repository.MemoRepository, repository.TagRepository) {
	return db.NewMemoRepository(infra.GetDBM()), db.NewTagRepository(infra.GetDBM())
}

// connectTestDB DB接続
func connectTestDB() {
	infra.ConnectTestDB()
}

// closeTestDB DB切断
func closeTestDB() {
	infra.CloseTestDB()
}
