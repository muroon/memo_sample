package di

import (
	"memo_sample/domain/repository"
	"memo_sample/adapter/memory"
	"memo_sample/adapter/db"
	"memo_sample/infra"
)

// InjectMemoInMemoryRepository inject repository
func InjectMemoInMemoryRepository() repository.MemoRepository {
	return &memory.MemoRepository{}
}

// InjectMemoDBRepository inject repository
func InjectMemoDBRepository(info *infra.InfraInfo) repository.MemoRepository {
	return &db.MemoRepository{DB: info.DB}
}