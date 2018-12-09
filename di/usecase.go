package di

import (
	"memo_sample/domain/repository"
	"memo_sample/usecase"
)

// InjectMemoUsecase inject memo usecase
func InjectMemoUsecase(
	transactionRepository repository.TransactionRepository,
	memoRepository repository.MemoRepository,
	tagagRepository repository.TagRepository) usecase.Memo {
	return usecase.NewMemo(transactionRepository, memoRepository, tagagRepository)
}

// InjectUsecaseIterator inject usecase itetator
func InjectUsecaseIterator(pre usecase.Presenter, memo usecase.Memo) usecase.Interactor {
	return usecase.NewInteractor(pre, memo)
}
