package di

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/memory"
	"memo_sample/domain/repository"
	"memo_sample/infra"
)

// InjectInMemoryRepository inject repository
func InjectInMemoryRepository() (repository.MemoRepository, repository.TagRepository) {
	return memory.NewMemoRepository(), memory.NewTagRepository()
}

// InjectDBRepository inject repository
func InjectDBRepository(info *infra.DbInfo) (repository.MemoRepository, repository.TagRepository) {
	return db.NewMemoRepository(info.DB), db.NewTagRepository(info.DB)
}
