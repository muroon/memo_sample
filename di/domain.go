package di

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/memory"
	"memo_sample/domain/repository"
	"memo_sample/infra"
)

// InjectMemoInMemoryRepository inject repository
func InjectMemoInMemoryRepository() repository.MemoRepository {
	return memory.NewMemoRepository()
}

// InjectMemoDBRepository inject repository
func InjectMemoDBRepository(info *infra.DbInfo) repository.MemoRepository {
	return db.NewMemoRepository(info.DB)
}
