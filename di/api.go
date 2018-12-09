package di

import (
	"memo_sample/interface/api"
	"memo_sample/usecase"
)

// InjectMemoryAPI inject api
func InjectMemoryAPI() api.API {
	iterator := InjectUsecaseIterator(
		InjectPresenter(),
		InjectMemoUsecase(InjectInMemoryRepository()))

	return api.NewAPI(
		iterator,
		InjectLog())
}

// InjectDBAPI inject api
func InjectDBAPI() api.API {
	iterator := InjectUsecaseIterator(
		InjectPresenter(),
		InjectMemoUsecase(InjectDBRepository()))

	return api.NewAPI(
		iterator,
		InjectLog())
}

// InjectPresenter inject presenter
func InjectPresenter() usecase.Presenter {
	return api.NewPresenter(InjectRender(), InjectLog())
}
