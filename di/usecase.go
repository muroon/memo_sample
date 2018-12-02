package di

import (
	"memo_sample/domain/repository"
	"memo_sample/usecase"
)

// InjectMemoUsecase inject memo usecase
func InjectMemoUsecase(memoRepository repository.MemoRepository) usecase.Memo {
	return usecase.NewMemo(memoRepository)
}
