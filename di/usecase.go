package di

import (
	"memo_sample/domain/repository"
	"memo_sample/usecase"
)

// InjectMemoUsecase inject memo usecase
func InjectMemoUsecase(memoRepo repository.MemoRepository, tagRepo repository.TagRepository) usecase.Memo {
	return usecase.NewMemo(memoRepo, tagRepo)
}
