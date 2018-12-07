package di

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/memory"
	"memo_sample/domain/repository"
	"memo_sample/infra"
)

// InjectInMemoryRepository inject repository
func InjectInMemoryRepository() (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return memory.NewTransactionRepository(), memory.NewMemoRepository(), memory.NewTagRepository()
}

// InjectDBRepository inject repository
func InjectDBRepository(dbm *infra.DBM) (repository.TransactionRepository, repository.MemoRepository, repository.TagRepository) {
	return db.NewTransactionRepository(dbm), db.NewMemoRepository(dbm), db.NewTagRepository(dbm)
}
