package di

import (
	"memo_sample/domain/repository"
	"memo_sample/adapter/memory"
)

// InjectMemoInMemoryRepository inject repository
func InjectMemoInMemoryRepository() repository.MemoRepository {
	return &memory.MemoRepository{}
}