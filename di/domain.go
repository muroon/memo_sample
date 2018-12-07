package di

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/memory"
	"memo_sample/domain/repository"
)

// InjectInMemoryRepository inject repository
func InjectInMemoryRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return memory.NewTransactionRepository(), memory.NewMemoRepository(), memory.NewTagRepository()
}

// InjectDBRepository inject repository
func InjectDBRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return db.NewTransactionRepository(), db.NewMemoRepository(), db.NewTagRepository()
}
